package java

import (
	"bytes"
	"log"

	"veyron2/vdl/compile"
)

const primitiveTmpl = `
// This file was auto-generated by the veyron vdl tool.
// Source: {{.Source}}
package {{.PackagePath}};

/**
 * type {{.Name}} {{.VdlTypeString}} {{.Doc}}
 **/
public final class {{.Name}} {
    private {{.BaseType}} value;

    public {{.Name}}({{.BaseType}} value) {
        this.value = value;
    }
    public {{.BaseType}} getValue() { return this.value; }

    public void setValue({{.BaseType}} value) { this.value = value; }

    @Override
    public boolean equals(java.lang.Object obj) {
        if (this == obj) return true;
        if (obj == null) return false;
        if (this.getClass() != obj.getClass()) return false;
        final {{.ObjectType}} other = ({{.ObjectType}})obj;
        {{ if .IsClass }}
        if (this.value == null) {
            return other.value == null;
        }
        return this.value.equals(other.value);
        {{ else }}
        return this.value == other.value;
        {{ end }}
    }
    @Override
    public int hashCode() {
        return {{.HashcodeComputation}};
    }
}
`

// genJavaPrimitiveFile generates the Java class file for the provided user-defined type.
func genJavaPrimitiveFile(tdef *compile.TypeDef, env *compile.Env) JavaFileInfo {
	data := struct {
		BaseType            string
		Doc                 string
		HashcodeComputation string
		IsClass             bool
		Name                string
		ObjectType          string
		PackagePath         string
		Source              string
		VdlTypeString       string
	}{
		BaseType:            javaType(tdef.BaseType, false, env),
		Doc:                 javaDocInComment(tdef.Doc),
		HashcodeComputation: javaHashCode("value", tdef.BaseType, env),
		IsClass:             isClass(tdef.BaseType, env),
		Name:                tdef.Name,
		ObjectType:          javaType(tdef.Type, true, env),
		PackagePath:         javaPath(javaGenPkgPath(tdef.File.Package.Path)),
		Source:              tdef.File.BaseName,
		VdlTypeString:       tdef.BaseType.String(),
	}
	var buf bytes.Buffer
	err := parseTmpl("primitive", primitiveTmpl).Execute(&buf, data)
	if err != nil {
		log.Fatalf("vdl: couldn't execute primitive template: %v", err)
	}
	return JavaFileInfo{
		Name: tdef.Name + ".java",
		Data: buf.Bytes(),
	}
}

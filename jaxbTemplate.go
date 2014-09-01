package main

type JaxbMainClassInfo struct {
	PackageName       string
	BaseXMLClassName  string
	SourceXMLFilename string
}

type JaxbClassInfo struct {
	Name                   string
	Root                   bool
	PackageName, ClassName string
	Attributes             []*JaxbAttribute
	Fields                 []*JaxbField
	HasValue               bool
	ValueType              string
}

type JaxbAttribute struct {
	Name      string
	NameUpper string
	NameLower string
	NameSpace string
}
type JaxbField struct {
	TypeName  string
	Name      string
	NameUpper string
	NameLower string
	NameSpace string
	Repeats   bool
}

func (jb *JaxbClassInfo) init() {
	jb.Attributes = make([]*JaxbAttribute, 0)
	jb.Fields = make([]*JaxbField, 0)
}

const jaxbClassTemplate = `
// Generated by chidley https://github.com/gnewton/chidley

package {{.PackageName}}.xml;

import java.util.ArrayList;
import javax.xml.bind.annotation.*;

@XmlAccessorType(XmlAccessType.FIELD)
@XmlRootElement(name="{{.Name}}")
public class {{.ClassName}} {
{{if .Attributes}}
    // Attributes{{end}}
{{range .Attributes}}
{{if .NameSpace}}    
@XmlAttribute(namespace = "{{.NameSpace}}"){{else}}    @XmlAttribute(name="{{.Name}}"){{end}}
    public String {{.NameLower}};{{end}}
{{if .Fields}}
    // Fields{{end}}{{range .Fields}}    
    @XmlElement(name="{{.Name}}")
    {{if .Repeats}}public ArrayList<{{.TypeName}}> {{.NameLower}}{{else}}public {{.TypeName}} {{.NameLower}}{{end}};
{{end}}
{{if .HasValue}}
    // Value
    @XmlValue
    public {{.ValueType}} tagValue;{{end}}
}
`

const jaxbMainTemplate = `
// Generated by chidley https://github.com/gnewton/chidley

package {{.PackageName}};
 
import java.io.File;
import javax.xml.bind.JAXBContext;
import javax.xml.bind.JAXBException;
import javax.xml.bind.Unmarshaller;
import {{.PackageName}}.xml.{{.BaseXMLClassName}};
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
 
public class Main {
	public static void main(String[] args) {
	 try {
		File file = new File("{{.SourceXMLFilename}}");
		JAXBContext jaxbContext = JAXBContext.newInstance({{.BaseXMLClassName}}.class);
 
		Unmarshaller jaxbUnmarshaller = jaxbContext.createUnmarshaller();
                {{.BaseXMLClassName}} root = ({{.BaseXMLClassName}}) jaxbUnmarshaller.unmarshal(file);

		Gson gson = new GsonBuilder().setPrettyPrinting().create();
		System.out.println(gson.toJson(root));

	  } catch (JAXBException e) {
		e.printStackTrace();
	  }
	}
}
`
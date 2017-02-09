//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		nvd.nist.gov/schema/scap-core_0.1.xsd
package go_ScapCore01

import (
	xml "github.com/linuxisnotunix/Vulnerobot/modules/models/xsd/www.w3.org/2009/01/xml.xsd_go"
	xsdt "github.com/metaleap/go-xsd/types"
)

//	Define the format for acceptable
//	searchableCPE Names.  The '*' may be used at the
//	beginning, end or, around, and will be interpreted as a
//	wildcard character.
type TcpeSearchableNamePatternType xsdt.AnyURI

//	Since TcpeSearchableNamePatternType is just a simple String type, this merely sets the current value from the specified string.
func (me *TcpeSearchableNamePatternType) Set(s string) { (*xsdt.AnyURI)(me).Set(s) }

//	Since TcpeSearchableNamePatternType is just a simple String type, this merely returns the current string value.
func (me TcpeSearchableNamePatternType) String() string { return xsdt.AnyURI(me).String() }

//	This convenience method just performs a simple type conversion to TcpeSearchableNamePatternType's alias type xsdt.AnyURI.
func (me TcpeSearchableNamePatternType) ToXsdtAnyURI() xsdt.AnyURI { return xsdt.AnyURI(me) }

type XsdGoPkgHasElem_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_ struct {
	CpeSearchableName TcpeSearchableNamePatternType `xml:"http://scap.nist.gov/schema/scap-core/0.1 cpe-searchable-name"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_ instance.
func (me *XsdGoPkgHasElem_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

//	Define the format for acceptable CPE Names. An urn format is used with the id starting with the word oval followed by a unique string, followed by the three letter code 'def', and ending with an integer.
type TcpeNamePatternType xsdt.AnyURI

//	Since TcpeNamePatternType is just a simple String type, this merely sets the current value from the specified string.
func (me *TcpeNamePatternType) Set(s string) { (*xsdt.AnyURI)(me).Set(s) }

//	Since TcpeNamePatternType is just a simple String type, this merely returns the current string value.
func (me TcpeNamePatternType) String() string { return xsdt.AnyURI(me).String() }

//	This convenience method just performs a simple type conversion to TcpeNamePatternType's alias type xsdt.AnyURI.
func (me TcpeNamePatternType) ToXsdtAnyURI() xsdt.AnyURI { return xsdt.AnyURI(me) }

type XsdGoPkgHasElem_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_ struct {
	CpeName TcpeNamePatternType `xml:"http://scap.nist.gov/schema/scap-core/0.1 cpe-name"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_ instance.
func (me *XsdGoPkgHasElem_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasGroup_CpeReferenceGroup struct {
	XsdGoPkgHasElem_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_

	XsdGoPkgHasElem_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_
}

//	If the WalkHandlers.XsdGoPkgHasGroup_CpeReferenceGroup function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasGroup_CpeReferenceGroup instance as the single argument. Then calls the Walk() method on 2/2 embed(s) and 0/0 field(s) belonging to this XsdGoPkgHasGroup_CpeReferenceGroup instance.
func (me *XsdGoPkgHasGroup_CpeReferenceGroup) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasGroup_CpeReferenceGroup; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.XsdGoPkgHasElem_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if err = me.XsdGoPkgHasElem_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasCdata struct {
	XsdGoPkgCDATA string `xml:",chardata"`
}

//	If the WalkHandlers.XsdGoPkgHasCdata function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasCdata instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasCdata instance.
func (me *XsdGoPkgHasCdata) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasCdata; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

//	Data type for the check element, a checking system specification URI, string content, and an optional external file reference. The checking system specification should be the URI for a particular version of OVAL or a related system testing language, and the content will be an identifier of a test written in that language. The external file reference could be used to point to the file in which the content test identifier is defined.
type XsdGoPkgHasAttr_System_XsdtAnyURI_ struct {
	System xsdt.AnyURI `xml:"http://scap.nist.gov/schema/scap-core/0.1 system,attr"`
}

type XsdGoPkgHasAttr_Href_XsdtAnyURI_ struct {
	Href xsdt.AnyURI `xml:"http://scap.nist.gov/schema/scap-core/0.1 href,attr"`
}

type XsdGoPkgHasAttr_Name_XsdtToken_ struct {
	Name xsdt.Token `xml:"http://scap.nist.gov/schema/scap-core/0.1 name,attr"`
}

type TcheckReferenceType struct {
	XsdGoPkgHasAttr_Href_XsdtAnyURI_

	XsdGoPkgHasAttr_Name_XsdtToken_

	XsdGoPkgHasAttr_System_XsdtAnyURI_
}

//	If the WalkHandlers.TcheckReferenceType function is not nil (ie. was set by outside code), calls it with this TcheckReferenceType instance as the single argument. Then calls the Walk() method on 0/3 embed(s) and 0/0 field(s) belonging to this TcheckReferenceType instance.
func (me *TcheckReferenceType) Walk() (err error) {
	if fn := WalkHandlers.TcheckReferenceType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

//	Type for a reference in the description of a CPE item. This would normally be used to point to extra descriptive material, or the supplier's web site, or the platform documentation. It consists of a piece of text (intended to be human-readable) and a URI (intended to be a URL, and point to a real resource).
//	This type allows the xml:lang attribute to associate a specific language with an element's string content.
type TtextType struct {
	XsdtString

	xml.XsdGoPkgHasAttr_Lang
}

//	If the WalkHandlers.TtextType function is not nil (ie. was set by outside code), calls it with this TtextType instance as the single argument. Then calls the Walk() method on 0/2 embed(s) and 0/0 field(s) belonging to this TtextType instance.
func (me *TtextType) Walk() (err error) {
	if fn := WalkHandlers.TtextType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type TreferenceType struct {
	TtextType

	XsdGoPkgHasAttr_Href_XsdtAnyURI_
}

//	If the WalkHandlers.TreferenceType function is not nil (ie. was set by outside code), calls it with this TreferenceType instance as the single argument. Then calls the Walk() method on 1/2 embed(s) and 0/0 field(s) belonging to this TreferenceType instance.
func (me *TreferenceType) Walk() (err error) {
	if fn := WalkHandlers.TreferenceType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TtextType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type TsearchableCpeReferencesType struct {
	XsdGoPkgHasGroup_CpeReferenceGroup
}

//	If the WalkHandlers.TsearchableCpeReferencesType function is not nil (ie. was set by outside code), calls it with this TsearchableCpeReferencesType instance as the single argument. Then calls the Walk() method on 1/1 embed(s) and 0/0 field(s) belonging to this TsearchableCpeReferencesType instance.
func (me *TsearchableCpeReferencesType) Walk() (err error) {
	if fn := WalkHandlers.TsearchableCpeReferencesType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.XsdGoPkgHasGroup_CpeReferenceGroup.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_ struct {
	CpeNames []TcpeNamePatternType `xml:"http://scap.nist.gov/schema/scap-core/0.1 cpe-name"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_ instance.
func (me *XsdGoPkgHasElems_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_NotesequencenotesTypeschema_Note_TtextType_ struct {
	Note *TtextType `xml:"http://scap.nist.gov/schema/scap-core/0.1 note"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_NotesequencenotesTypeschema_Note_TtextType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_NotesequencenotesTypeschema_Note_TtextType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElem_NotesequencenotesTypeschema_Note_TtextType_ instance.
func (me *XsdGoPkgHasElem_NotesequencenotesTypeschema_Note_TtextType_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_NotesequencenotesTypeschema_Note_TtextType_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.Note.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasAttr_Value_XsdtToken_ struct {
	Value xsdt.Token `xml:"http://scap.nist.gov/schema/scap-core/0.1 value,attr"`
}

//	The name pattern of a CPE component.
type TcpeComponentPatternType xsdt.Token

//	Since TcpeComponentPatternType is just a simple String type, this merely sets the current value from the specified string.
func (me *TcpeComponentPatternType) Set(s string) { (*xsdt.Token)(me).Set(s) }

//	Since TcpeComponentPatternType is just a simple String type, this merely returns the current string value.
func (me TcpeComponentPatternType) String() string { return xsdt.Token(me).String() }

//	This convenience method just performs a simple type conversion to TcpeComponentPatternType's alias type xsdt.Token.
func (me TcpeComponentPatternType) ToXsdtToken() xsdt.Token { return xsdt.Token(me) }

//	The name pattern of the CPE part component.
type TcpePartComponentPatternType TcpeComponentPatternType

//	Since TcpePartComponentPatternType is just a simple String type, this merely sets the current value from the specified string.
func (me *TcpePartComponentPatternType) Set(s string) { (*TcpeComponentPatternType)(me).Set(s) }

//	Since TcpePartComponentPatternType is just a simple String type, this merely returns the current string value.
func (me TcpePartComponentPatternType) String() string { return TcpeComponentPatternType(me).String() }

//	This convenience method just performs a simple type conversion to TcpePartComponentPatternType's alias type TcpeComponentPatternType.
func (me TcpePartComponentPatternType) ToTcpeComponentPatternType() TcpeComponentPatternType {
	return TcpeComponentPatternType(me)
}

type TcweNamePatternType xsdt.Token

//	Since TcweNamePatternType is just a simple String type, this merely sets the current value from the specified string.
func (me *TcweNamePatternType) Set(s string) { (*xsdt.Token)(me).Set(s) }

//	Since TcweNamePatternType is just a simple String type, this merely returns the current string value.
func (me TcweNamePatternType) String() string { return xsdt.Token(me).String() }

//	This convenience method just performs a simple type conversion to TcweNamePatternType's alias type xsdt.Token.
func (me TcweNamePatternType) ToXsdtToken() xsdt.Token { return xsdt.Token(me) }

type TcheckSearchType struct {
	XsdGoPkgHasAttr_System_XsdtAnyURI_

	XsdGoPkgHasAttr_Name_XsdtToken_
}

//	If the WalkHandlers.TcheckSearchType function is not nil (ie. was set by outside code), calls it with this TcheckSearchType instance as the single argument. Then calls the Walk() method on 0/2 embed(s) and 0/0 field(s) belonging to this TcheckSearchType instance.
func (me *TcheckSearchType) Walk() (err error) {
	if fn := WalkHandlers.TcheckSearchType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_NotesequencenotesTypeschema_Note_TtextType_ struct {
	Notes []*TtextType `xml:"http://scap.nist.gov/schema/scap-core/0.1 note"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_NotesequencenotesTypeschema_Note_TtextType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_NotesequencenotesTypeschema_Note_TtextType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_NotesequencenotesTypeschema_Note_TtextType_ instance.
func (me *XsdGoPkgHasElems_NotesequencenotesTypeschema_Note_TtextType_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_NotesequencenotesTypeschema_Note_TtextType_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.Notes {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

//	The notesType defines an element that consists of one or more child note elements. It is assumed that each of these note elements are representative of the same language as defined by their parent.
type TnotesType struct {
	XsdGoPkgHasElems_NotesequencenotesTypeschema_Note_TtextType_
}

//	If the WalkHandlers.TnotesType function is not nil (ie. was set by outside code), calls it with this TnotesType instance as the single argument. Then calls the Walk() method on 1/1 embed(s) and 0/0 field(s) belonging to this TnotesType instance.
func (me *TnotesType) Walk() (err error) {
	if fn := WalkHandlers.TnotesType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.XsdGoPkgHasElems_NotesequencenotesTypeschema_Note_TtextType_.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type TtagType struct {
	XsdGoPkgHasAttr_Value_XsdtToken_

	XsdGoPkgHasAttr_Name_XsdtToken_
}

//	If the WalkHandlers.TtagType function is not nil (ie. was set by outside code), calls it with this TtagType instance as the single argument. Then calls the Walk() method on 0/2 embed(s) and 0/0 field(s) belonging to this TtagType instance.
func (me *TtagType) Walk() (err error) {
	if fn := WalkHandlers.TtagType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_ struct {
	CpeSearchableNames []TcpeSearchableNamePatternType `xml:"http://scap.nist.gov/schema/scap-core/0.1 cpe-searchable-name"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_ instance.
func (me *XsdGoPkgHasElems_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

var (
	//	Set this to false to break a Walk() immediately as soon as the first error is returned by a custom handler function.
	//	If true, Walk() proceeds and accumulates all errors in the WalkErrors slice.
	WalkContinueOnError = true
	//	Contains all errors accumulated during Walk()s. If you're using this, you need to reset this yourself as needed prior to a fresh Walk().
	WalkErrors []error
	//	Your custom error-handling function, if required.
	WalkOnError func(error)
	//	Provides 15 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
	//	If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
	WalkHandlers = &XsdGoPkgWalkHandlers{}
)

//	Provides 15 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
//	If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
type XsdGoPkgWalkHandlers struct {
	TsearchableCpeReferencesType                                                                                     func(*TsearchableCpeReferencesType, bool) error
	XsdGoPkgHasElem_NotesequencenotesTypeschema_Note_TtextType_                                                      func(*XsdGoPkgHasElem_NotesequencenotesTypeschema_Note_TtextType_, bool) error
	TnotesType                                                                                                       func(*TnotesType, bool) error
	XsdGoPkgHasElems_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_ func(*XsdGoPkgHasElems_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_, bool) error
	TtagType                                                                                                         func(*TtagType, bool) error
	TreferenceType                                                                                                   func(*TreferenceType, bool) error
	XsdGoPkgHasElems_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_                               func(*XsdGoPkgHasElems_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_, bool) error
	XsdGoPkgHasElem_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_  func(*XsdGoPkgHasElem_CpeSearchableNamechoicecpeReferenceGroupschema_CpeSearchableName_TcpeSearchableNamePatternType_, bool) error
	XsdGoPkgHasElem_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_                                func(*XsdGoPkgHasElem_CpeNamechoicecpeReferenceGroupschema_CpeName_TcpeNamePatternType_, bool) error
	XsdGoPkgHasGroup_CpeReferenceGroup                                                                               func(*XsdGoPkgHasGroup_CpeReferenceGroup, bool) error
	TcheckReferenceType                                                                                              func(*TcheckReferenceType, bool) error
	TtextType                                                                                                        func(*TtextType, bool) error
	XsdGoPkgHasCdata                                                                                                 func(*XsdGoPkgHasCdata, bool) error
	TcheckSearchType                                                                                                 func(*TcheckSearchType, bool) error
	XsdGoPkgHasElems_NotesequencenotesTypeschema_Note_TtextType_                                                     func(*XsdGoPkgHasElems_NotesequencenotesTypeschema_Note_TtextType_, bool) error
}
type XsdtString xsdt.String

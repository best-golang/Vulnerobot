//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		nvd.nist.gov/schema/cve_0.1.1.xsd
package go_Cve011

import (
	scap_core "github.com/linuxisnotunix/Vulnerobot/modules/models/xsd/nvd.nist.gov/schema/scap-core_0.1.xsd_go"
	xsdt "github.com/metaleap/go-xsd/types"
)

//	Discretionary information and links relevant to a given vulnerability referenced by the CVE
type XsdGoPkgHasElems_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_ struct {
	//	Discretionary information and links relevant to a given vulnerability referenced by the CVE
	Referenceses []scap_core.TreferenceType `xml:"http://scap.nist.gov/schema/cve/0.1 references"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_ instance.
func (me *XsdGoPkgHasElems_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_; me != nil {
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

//	Format for CVE Names is CVE-YYYY-NNNN+, where YYYY is the year of publication and NNNN is a sequence number.
type TcveNamePatternType xsdt.Token

//	Since TcveNamePatternType is just a simple String type, this merely returns the current string value.
func (me TcveNamePatternType) String() string { return xsdt.Token(me).String() }

//	This convenience method just performs a simple type conversion to TcveNamePatternType's alias type xsdt.Token.
func (me TcveNamePatternType) ToXsdtToken() xsdt.Token { return xsdt.Token(me) }

//	Since TcveNamePatternType is just a simple String type, this merely sets the current value from the specified string.
func (me *TcveNamePatternType) Set(s string) { (*xsdt.Token)(me).Set(s) }

//	CVE name in the CVE-YYYY-NNNN+ format
type XsdGoPkgHasAttr_Id_TcveNamePatternType_ struct {
	//	CVE name in the CVE-YYYY-NNNN+ format
	Id TcveNamePatternType `xml:"http://scap.nist.gov/schema/cve/0.1 id,attr"`
}

//	Status of Vulnerability -- Candidate, Entry, Deprecated
//	Enumeration containing valid values for CVE status: Candidate, Entry, and Deprecated
type TcveStatus xsdt.Token

//	Returns true if the value of this enumerated TcveStatus is "CANDIDATE".
func (me TcveStatus) IsCandidate() bool { return me.String() == "CANDIDATE" }

//	Returns true if the value of this enumerated TcveStatus is "ENTRY".
func (me TcveStatus) IsEntry() bool { return me.String() == "ENTRY" }

//	Returns true if the value of this enumerated TcveStatus is "DEPRECATED".
func (me TcveStatus) IsDeprecated() bool { return me.String() == "DEPRECATED" }

//	Since TcveStatus is just a simple String type, this merely sets the current value from the specified string.
func (me *TcveStatus) Set(s string) { (*xsdt.Token)(me).Set(s) }

//	Since TcveStatus is just a simple String type, this merely returns the current string value.
func (me TcveStatus) String() string { return xsdt.Token(me).String() }

//	This convenience method just performs a simple type conversion to TcveStatus's alias type xsdt.Token.
func (me TcveStatus) ToXsdtToken() xsdt.Token { return xsdt.Token(me) }

type XsdGoPkgHasElem_StatussequencecveTypeschema_Status_TcveStatus_ struct {
	//	Status of Vulnerability -- Candidate, Entry, Deprecated
	Status TcveStatus `xml:"http://scap.nist.gov/schema/cve/0.1 status"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_StatussequencecveTypeschema_Status_TcveStatus_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_StatussequencecveTypeschema_Status_TcveStatus_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_StatussequencecveTypeschema_Status_TcveStatus_ instance.
func (me *XsdGoPkgHasElem_StatussequencecveTypeschema_Status_TcveStatus_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_StatussequencecveTypeschema_Status_TcveStatus_; me != nil {
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

//	Free text field to describe the vulnerability
type XsdGoPkgHasElems_DescriptionsequencecveTypeschema_Description_XsdtString_ struct {
	//	Free text field to describe the vulnerability
	Descriptions []xsdt.String `xml:"http://scap.nist.gov/schema/cve/0.1 description"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_DescriptionsequencecveTypeschema_Description_XsdtString_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_DescriptionsequencecveTypeschema_Description_XsdtString_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_DescriptionsequencecveTypeschema_Description_XsdtString_ instance.
func (me *XsdGoPkgHasElems_DescriptionsequencecveTypeschema_Description_XsdtString_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_DescriptionsequencecveTypeschema_Description_XsdtString_; me != nil {
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

//	Discretionary information and links relevant to a given vulnerability referenced by the CVE
type XsdGoPkgHasElem_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_ struct {
	//	Discretionary information and links relevant to a given vulnerability referenced by the CVE
	References scap_core.TreferenceType `xml:"http://scap.nist.gov/schema/cve/0.1 references"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_ instance.
func (me *XsdGoPkgHasElem_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_; me != nil {
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

//	Status of Vulnerability -- Candidate, Entry, Deprecated
type XsdGoPkgHasElems_StatussequencecveTypeschema_Status_TcveStatus_ struct {
	//	Status of Vulnerability -- Candidate, Entry, Deprecated
	Statuses []TcveStatus `xml:"http://scap.nist.gov/schema/cve/0.1 status"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_StatussequencecveTypeschema_Status_TcveStatus_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_StatussequencecveTypeschema_Status_TcveStatus_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_StatussequencecveTypeschema_Status_TcveStatus_ instance.
func (me *XsdGoPkgHasElems_StatussequencecveTypeschema_Status_TcveStatus_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_StatussequencecveTypeschema_Status_TcveStatus_; me != nil {
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

//	Free text field to describe the vulnerability
type XsdGoPkgHasElem_DescriptionsequencecveTypeschema_Description_XsdtString_ struct {
	//	Free text field to describe the vulnerability
	Description xsdt.String `xml:"http://scap.nist.gov/schema/cve/0.1 description"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_DescriptionsequencecveTypeschema_Description_XsdtString_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_DescriptionsequencecveTypeschema_Description_XsdtString_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_DescriptionsequencecveTypeschema_Description_XsdtString_ instance.
func (me *XsdGoPkgHasElem_DescriptionsequencecveTypeschema_Description_XsdtString_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_DescriptionsequencecveTypeschema_Description_XsdtString_; me != nil {
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

type TcveType struct {
	//	Status of Vulnerability -- Candidate, Entry, Deprecated
	XsdGoPkgHasElem_StatussequencecveTypeschema_Status_TcveStatus_

	//	Free text field to describe the vulnerability
	XsdGoPkgHasElem_DescriptionsequencecveTypeschema_Description_XsdtString_

	//	Discretionary information and links relevant to a given vulnerability referenced by the CVE
	XsdGoPkgHasElems_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_

	//	CVE name in the CVE-YYYY-NNNN+ format
	XsdGoPkgHasAttr_Id_TcveNamePatternType_
}

//	If the WalkHandlers.TcveType function is not nil (ie. was set by outside code), calls it with this TcveType instance as the single argument. Then calls the Walk() method on 3/4 embed(s) and 0/0 field(s) belonging to this TcveType instance.
func (me *TcveType) Walk() (err error) {
	if fn := WalkHandlers.TcveType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.XsdGoPkgHasElem_StatussequencecveTypeschema_Status_TcveStatus_.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if err = me.XsdGoPkgHasElem_DescriptionsequencecveTypeschema_Description_XsdtString_.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if err = me.XsdGoPkgHasElems_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
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

var (
	//	Set this to false to break a Walk() immediately as soon as the first error is returned by a custom handler function.
	//	If true, Walk() proceeds and accumulates all errors in the WalkErrors slice.
	WalkContinueOnError = true
	//	Contains all errors accumulated during Walk()s. If you're using this, you need to reset this yourself as needed prior to a fresh Walk().
	WalkErrors []error
	//	Your custom error-handling function, if required.
	WalkOnError func(error)
	//	Provides 8 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
	//	If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
	WalkHandlers = &XsdGoPkgWalkHandlers{}
)

//	Provides 8 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
//	If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
type XsdGoPkgWalkHandlers struct {
	XsdGoPkgHasCdata                                                                    func(*XsdGoPkgHasCdata, bool) error
	XsdGoPkgHasElems_StatussequencecveTypeschema_Status_TcveStatus_                     func(*XsdGoPkgHasElems_StatussequencecveTypeschema_Status_TcveStatus_, bool) error
	XsdGoPkgHasElem_DescriptionsequencecveTypeschema_Description_XsdtString_            func(*XsdGoPkgHasElem_DescriptionsequencecveTypeschema_Description_XsdtString_, bool) error
	TcveType                                                                            func(*TcveType, bool) error
	XsdGoPkgHasElems_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_ func(*XsdGoPkgHasElems_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_, bool) error
	XsdGoPkgHasElem_StatussequencecveTypeschema_Status_TcveStatus_                      func(*XsdGoPkgHasElem_StatussequencecveTypeschema_Status_TcveStatus_, bool) error
	XsdGoPkgHasElems_DescriptionsequencecveTypeschema_Description_XsdtString_           func(*XsdGoPkgHasElems_DescriptionsequencecveTypeschema_Description_XsdtString_, bool) error
	XsdGoPkgHasElem_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_  func(*XsdGoPkgHasElem_ReferencessequencecveTypeschema_References_ScapCoreTreferenceType_, bool) error
}

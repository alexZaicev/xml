parser grammar XMLParser;

options { tokenVocab = XMLLexer; }

document    :   prolog? misc* element misc*;

prolog      :   XMLDeclOpen attribute* SPECIAL_CLOSE ;

content     :   chardata?
                ((element | reference | CDATA | PI | COMMENT) chardata?)* ;

element     :   '<' beginning=Name attribute* '>' content '<' '/' ending=Name '>'
            |   '<' beginning=Name attribute* '/>'
            ;

reference   :   EntityRef | CharRef ;

attribute   :   name=Name '=' value=STRING ; // Our STRING is AttValue in specname

/** ``All text that is not markup constitutes the character data of
 *  the document.''
 */
chardata    :   value=TEXT | SEA_WS ;

misc        :   COMMENT | PI | SEA_WS ;
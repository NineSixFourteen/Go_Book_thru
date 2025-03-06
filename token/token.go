package token   

type TokenType string

type Token struct {
    Type TokenType
    Literal string 
}

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"
    //Ident
    IDENT   = "IDENT" 
    INT     = "INT"

    // OPs
    ASSIGN    = "="
    PLUS      = "+"
    MINUS     = "-"
    ASTERISK  = "*"
    SLASH     = "/"
    // Del 
    COMMA     = ","
    SEMICOLON = ";"

    LPAREN  = "("
    RPAREN  = ")"
    LBRACE  = "{"
    RBRACE  = "}"

    //Keywords
    FUNCTION = "FUNCTION"
    LET      = "LET"
    RETURN   = "RETURN"
    IF       = "IF"
    ELSE     = "ELSE"

    //Logic 
    LT     = "<"
    GT     = ">"
    BANG   = "!"
    EQ     = "=="
    NOT_EQ = "!="

    //Bools
    TRUE = "True"
    FALSE = "False"


    
    

)

var keywords = map[string]TokenType{
    "fn": FUNCTION,
    "let": LET,
    "if": IF, 
    "else": ELSE ,
    "true": TRUE ,
    "false": FALSE,
    "return": RETURN,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok;
    }
    return IDENT;
}

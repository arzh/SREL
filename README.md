TODO: Rewrite this doc

SRXL
====

Structured Regular eXpression Language

Idea:
  Create an SQL like query that creates a regex
    Example: SRXL(`startline then "http" maybe "s" then "://" maybe "www." anythingbut " " eol`)
    
	Simple tokening:
```go

	tokens := strings.Split(q, " ")
	
	toCnt := false
	cnt := ""
  	for i, e := range tokens {
  		if e[0] == '"' && (e[len(e)-1] != '"' || len(e) <= 1) && !toCnt {
  			toCnt = true
  			cnt += e + " "
  			//fmt.Println("Need to continue,", cnt)
  			continue
  		}
  		
  		if toCnt { 
  			//fmt.Println("Continuing,", cnt, e)
  			if e[len(e)-1] == '"' {
  				e = cnt + e
  				cnt = ""
  				toCnt = false
  			} else {
  				cnt += e + " "
  				continue
  			}
  		}
  			
  		fmt.Println("token", i, " =", e)
  	}
```

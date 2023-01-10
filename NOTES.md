# Notes

## questions.json
 - The description is split into two lines, because all the 
results have html line breaks "<br/>". I would rather not inject html code, because it might pose a security risk.
 - The ids are same, than the ids of the URL result of the psychologies page. For this kind of public data, uuid is better.


 ## commands

 docker network create -d bridge db-network
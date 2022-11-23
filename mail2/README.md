## mail2

### Init
execute `mail2 init` to set up configures, configures are in  `~/.config/mail2` . And then edit this file.

### Usage

Send mail which subject is "test" to some email.

```shell
mail2 -s test -m mail_content -a /home/laily/aa.jpg aa@aa.com

mail2 -h show help
```

### TODO 
- Use pipe command
  echo “mail content”| mail2 -s test aa@aa.com  

- Read file as mail content
  mail2 -s test aa@aa.com< file  

- Send to multiple users  

  mail2 -s test -c aa@aa.com  bb@aa.com < file  


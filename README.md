# NiCloud

> Brute force public AWS, GCS, and DigitalOcean cloud services. 

- Download the default wordlist file [here](https://raw.githubusercontent.com/i5nipe/nicloud/master/files/cloud.txt).

- You don't need to specify the wordlist file if you put it in `~/.nipe/cloud.txt`.

- If none of the arguments are specified it will brute force *all*. *[-aws -gcs -dos]*

## ☕ Install

```bash
{your package manager} install awscli
```

```bash
go get github.com/i5nipe/nicloud
```


## ☕ Examples
```
nicloud -d {COMPANY} -w permlist.txt

nicloud -aws -d {COMPANY}

nicloud -dos -gcs -d {COMPANY}
```

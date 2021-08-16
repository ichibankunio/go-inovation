set -e

name=innovation2007
app_name=Innovation2007.app

cd bin

zip ${name}_linux_386.zip ${name}_linux_386
zip ${name}_linux_amd64.zip ${name}_linux_amd64
zip -r ${name}_darwin_amd64.zip ${app_name}
zip ${name}_windows_386.zip ${name}_windows_386.exe
zip ${name}_windows_amd64.zip ${name}_windows_amd64.exe

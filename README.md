
# Polish invoice creation system

## Goals

* Generate beautiful pdf invoices
* Respect all the rules defined by law
* Manage data as plain text
* Command line interface for operations
* Docker image containing all the prerequisites

## Command line interface

This application should conform to Unix philosophy.

### Available commands

#### Add party

```
AME:
   faktura party add - Add new party

USAGE:
   faktura party add [command options] [arguments...]

OPTIONS:
   --code value      Code of the party
   --name value      Name of the party
   --address value   Address of the party
   --address2 value  Second address line of the party
   --nip value       NIP of the party
   --regon value     Regon of the party
```

#### Modify party

```
NAME:
   faktura party modify - Modify party

USAGE:
   faktura party modify [command options] [arguments...]

OPTIONS:
   --code value               Code of the party
   --name value               Name of the party
   --address value            Address of the party
   --address2 value           Second address line of the party
   --nip value                NIP of the party
   --regon value              Regon of the party
   --numbering-pattern value  Invoice numbering pattern
   --active-from value        Start date from which this version is actual
```

#### Configure

* [ ] Set default party

#### Add invoice

```
NAME:
   faktura invoice add - Add new invoice

USAGE:
   faktura invoice add [command options] [arguments...]

OPTIONS:
   --buyer value           Code of the buyer
   --seller value          Name of the seller
   --issuanceDate value    Date of issuance
   --issuancePlace value   Place of issuance
   --sellDate value        Sell date
   --positionFormat value  Defines input format of the positions
   --positions value       Positions declared on invoice
```

#### Modify invoice

TODO

#### Delete invoice

TODO

#### Print invoice

```
NAME:
   faktura generate invoice - Print invoice

USAGE:
   faktura generate invoice [command options] [arguments...]

OPTIONS:
   --party value  Code of the party
   --last         Generate last invoice created for the party``
```

## Docker image content

Final docker image should contain application itself
as well as latex toolchain for generating pdf files
from tex source files.

## Building application

Automated build script is defined:

```sh
./build.sh
```

Simple rebuild and exec loop

```sh
sudo rm -f bin/* && ./build.sh && ./bin/faktura generate invoice --party [party_code] --last
```

## Tasks

* [ ] Enable golang modules
* [ ] Include cli, file processing, structured text parsing libraries
* [ ] Define initial structure for data
* [ ] Create initial template for invoice
* [ ] Understand correct role of git
* [ ] Creating docker image
    * [ ] Include Latex toolchain
    * [ ] Mount point for data

## External sources

* https://jakub.nadolny.info/faktura.tex
* https://ljvmiranda921.github.io/notebook/2018/04/23/postmortem-shift-to-docker/

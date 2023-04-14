# DNS-server-searcher
___________________________________

### Утилита командной строки, позволяющая работать с DNS протоколом, что позволяет осуществлять поиск серверов доменных имен, данные о которых хранятся в NS-записях доменов, позволяет осуществлять поиск почтовых серверов, данные о которых хранятся в MX-записях доменов и позволяет получить на основе IP адресов имена доменов, связанных с данными именами, и на основе имен доменов позволяет получить IP адреса связаныне с данными доменами.

* -d - осуществляет поиск доменных имен, связанных с переданными IP адресами формата IPv4 и IPv6
* -i - осуществляет поиск IP адресов формата IPv4 и IPv6, связанных с переданными доменными именами
* -ns - осуществляет поиск серверов доменных имен, хранящихся в NS-записях переданных доменов
* -mx - осуществляет поиск почтовых серверов, хранящихся в MX-записях переданных доменов
* -input - флаг предназначенный для указаний файла из которого будет осуществляться ввод данных 
* -output - флаг предназначенный для указаний файла в который будут выводиться данные

### Поиск доменных имен
-d - поиск **доменных имен** на основе переданных IP адресов

## Пример использования флага -d:
![result1](https://github.com/ellofae/DNS-server-searcher/blob/main/img/Screenshot%20from%202023-04-14%2018-07-42.png?raw=true)

### Поиск IP адресов
-i - поиск **IP адресов** на основе переданных доменов

## Пример использования флага -i:
![result2](https://github.com/ellofae/DNS-server-searcher/blob/main/img/Screenshot%20from%202023-04-14%2018-03-20.png?raw=true)

### Поиск серверов доменных имен
-ns - поиск серверов доменных имен на основе переданных доменов

## Пример использования флага -ns:
![result3](https://github.com/ellofae/DNS-server-searcher/blob/main/img/Screenshot%20from%202023-04-14%2018-40-05.png?raw=true)

### Поиск почтовых серверов
-mx - поиск почтовых серверов на основе переданных доменов

## Пример использования флага -mx:
![result4](https://github.com/ellofae/DNS-server-searcher/blob/main/img/Screenshot%20from%202023-04-14%2019-02-32.png?raw=true)

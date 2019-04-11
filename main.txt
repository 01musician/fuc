1. 十进制与十六进制互换
a: 使用bc：
echo "obase=16; 34" | bc
b：使用printf
printf "%x\n" 34
2.查看IP地址
Ifconfig -a

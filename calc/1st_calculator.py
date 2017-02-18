'''
basic calculator for two numbers. 
'''

from decimal import Decimal
class calculation():
    ''' class for actual calculation'''
    def __init__(self,number1,number2,operator):
        self.num1 = number1
        self.num2 = number2
        self.oper = operator
    def calculation(self):
        try:
            if(self.oper == '+'):
                return self.num1 + self.num2
            elif(self.oper == '-'):
                return self.num1 - self.num2
            elif(self.oper == '*'):
                return self.num1 * self.num2
            elif(self.oper == '/'):
                return float(self.num1) / float(self.num2)
            
        except Exception as ex:
            print "error: ",ex
            return "wrong"

        
def calculator_main(calculate_input,exp):
    ''' parsing the input string and generate tokens'''
    operators = ['*','+','-','/']
    numbers,exp = "",[]
    
    for i in calculate_input:
        if(i.isalpha() == True):
            if(calculate_input == "exit"):
                exit(0)
            print("invalid input")
            return "wrong"
        if(i.isdigit() == True or i == '.'):
            numbers += i
        elif(i in operators):
            if( numbers != ''):
                exp.append(numbers)
            exp.append(i)
            numbers = ""
    if(numbers != ''):
        exp.append(numbers)
    return exp

def main():
    flg = 0
    operators = ['*','+','-','/']
    while True:
        calculate_input = raw_input()
        exp = calculator_main(calculate_input,[])
        try :
            if(exp[0].isdigit() == True or '.' in exp[0] ):
                flg = 0
                if('.' in exp[0]):
                    exp[0] = Decimal(exp[0])
                    exp[2] = Decimal(exp[2])
                elif(exp[0].isdigit() == True):
                    exp[0] = int(exp[0])
                    exp[2] = int(exp[2])
            elif('.' in exp[1]):
                exp[1] = Decimal(exp[1])
            else:
                exp[1] = int(exp[1])
            if(flg == 0):
                num1 = exp[0]
                operator = exp[1]
                num2 = exp[2]
            else:
                num1 = res
                operator = exp[0]
                num2 = exp[1]
            calc = calculation(num1,num2,operator)
            res = calc.calculation()
            if(res == "wrong"):
                print "give Valid Input..."
                flg = 0
            elif(res or res == 0):
                print res
                flg = 1
        except Exception as ex:
            print "input Error",ex
            flg = 0
main()

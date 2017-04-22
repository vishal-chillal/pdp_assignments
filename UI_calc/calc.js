$(document).ready(function (){
	dot_count = 0;	
	$("input").keypress(function(e){
    var charCode = !e.charCode ? e.which : e.charCode;
    //console.log(charCode);
    if(charCode == 13){
    	$('.submitButton').click();
    }
    if(check_keys(charCode,dot_count) == false){
			e.preventDefault();
		}
	})
	$('.submitButton').click(function(e){
		ans = $('#ans').val();
		main_function(ans);
	});
});
function check_keys(charCode){
	if(charCode != 44 && charCode <= 57 && charCode >=42 || charCode == 0 || charCode == 8 || charCode == 114 || charCode == 94 ) {
			return true;
		}
	else{
		return false;
	}
}

function main_function(ans){
	expression = []
	if($('#op1').val()){
	expression.push($('#op1').val());
	}
	else{
	expression.push(ans);
	}
	expression.push($('#op2').val());
	expression.push($('#op').val());
	ans = calculate(expression);
	if(ans != ""){
		document.getElementById('ans').value = ans;
		document.getElementById('op1').value = ans;
	}
	else{
		ans = "please enter valid input";
		document.getElementById('ans').value = ans;
		document.getElementById('op1').value = 0;
	}
	
}
function calculate(expression){
	try {
	    
			if(expression[2] == '*'){
				ans = expression[0]*expression[1];
			}
			else if(expression[2] == '+'){
				ans = Number(expression[0])+ Number(expression[1]);
			}
			else if(expression[2] == '-'){
				ans = expression[0] - expression[1];
			}
			else if(expression[2] == '^'){
				ans = expression[0]^expression[1];
			}

			else if(expression[2] == '/'){
				if(expression[1] != 0){
					ans = expression[0]/expression[1];
				}
				else{
					ans = "";
				}
			}else{
				return "";
			}
		}
	catch(err) {
	    return "";
	}
	return ans;
}
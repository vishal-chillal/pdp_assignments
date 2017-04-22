$(document).ready(function (){
//$("#dateID").datepicker();
//Ajax: function(type, url, datatype, data, loadMessage, errorMessage,contentType, callBack, obj){
    $( "form" ).on( "submit", function(event){
        event.preventDefault();
        var log = $( "form").serializeArray()
        console.log("before clean json:\n",JSON.stringify(log));
        var clean_json = {};
        for(var j in log) {
            clean_json[log[j]["name"]] = log[j]["value"];
        }
        console.log("after clean json:\n",clean_json);
        var requestJson = {
            "Api" : "login",
            "Data" : JSON.stringify(clean_json).replace(/"/g,'\"');
        }
        var url = "http://localhost:81/apiserver?query="+requestJson;
        pdp.callWithAjax(
            "GET",
            url,
            "json",
            null,
            "Please Wait,kiti ghai?????",
            "Error occure while posting",
            "application/json",
            null,
            this
            );
    });
});


var pdp = {
    callWithAjax: function(type, url, datatype, data, loadMessage, errorMessage,contentType, callBack, obj){
     headers = {}
     var self = this;
     $.ajax({
         type: type,
         url: url,
         datatype: datatype,
         data: data,
         contentType: contentType,
         headers: headers,
         beforeSend: function(){
          self.alertFunction("beforeSEND");
      },
      error:function(err){
          self.alertFunction("error");
          console.log(err);
      },
      success:function(response){
		self.rerender(obj,response);
		self.alertFunction("success");
 },
 complete:function(xhr){
  self.alertFunction("complete");
}

});
 },

 alertFunction: function(func){
     console.log("call:"+func);
 }


    rerender: function(form_fields,response){
 	for (i in response) {
 	    form_fields[i].value = response[i];
	}
   }

};

function validateEmail(emailField){
    var reg = /^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;
    if (reg.test(emailField.value) == false)
    {
     alert('Invalid Email Address');
     return false;
 }
}

function validateMobile(mobField){
    var reg = /^([0-9]{10})/;
    if (reg.test(mobField.value) == false)
    {
     alert('Invalid Mobile Number');
     return false;
 }
}

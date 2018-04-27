
$(function() {
	 $('#reused_form').submit(function(e) {
		  e.preventDefault();
		  $form = $(this);
		  //show some response on the button
		  $('button[type="submit"]', $form).each(function()
				{
					 $btn = $(this);
					 $btn.prop('type','button' );
					 $btn.prop('orig_label',$btn.text());
					 $btn.text('Sending ...');
				});

		  $.ajax({
				type: "POST",
				url: '/posts',
				data: $form.serialize(),
				dataType: 'json',
				success: function (response) {
					 if (response.message == "success") {
						  alert("Success, save data !");
						  window.location.reload();
					 }else{
						  alert(response);
					 }
				},
		  });
	 });
});

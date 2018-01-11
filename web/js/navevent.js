$(document).ready(function(){
	var contact = $("#newContact")
	$("#addCont").click(function(){
		$(contact).show();
		$(contact).find("table input[type=text]").first().focus();
	});

	/* Hide New Contact form*/
	$("#form-cancel").click(function(){
	 	$(contact).hide();
	});

	/* Reset's the search and shows all content*/
	$("#reset-search").click(function() {
		$("#filter").val('');
		var $rows = $('#contact-details table tr');
		notificationTextSet();
		$rows.fadeIn();
	});

	/* Toggel del-selected class on span. */
	$(".del-def").click(function(){
		$(this).toggleClass('del-selected');
		$(this).parent().siblings().first().toggleClass('row-fcol-select');
	});

	function notificationTextSet(msg){
		if(msg !== undefined){
			$("#notification").text(msg);
		}else {
			$("#notification").text(':-)');
		}
	}

	// Deleting contacts
	$("#contact-details").submit(function(e){
		e.preventDefault();
		// this tells the server-side process that Ajax was used

		// returns string which is url Query part
		var	arr = [];
		var elements = [];
		var dataT = function() {
			var record = '',
				deleteQuery = '?',
				deleteString = "delete=",
				tmp = '';
			$(".del-selected").each(function(){
				tmp = $(this).attr('recordno');
				arr.push(parseFloat(tmp));
				elements.push($(this));
			});

			function pro(element, index, array){
				deleteQuery = deleteQuery + (deleteString + element + '&');
			}

			arr.forEach(pro);
			//console.log(deleteQuery);
			return deleteQuery.slice(0, deleteQuery.length - 1);
		}

		$.ajax ({
			type: 'GET',
			url: 'http://localhost:3000/delete'+ dataT(),
			success: function(data) {
				var rowdata = data.slice(1, data.length-1);
				var datastring = rowdata.split(" ");
				var datax = datastring.map(parseFloat);

				var idx = -1,
					delContactCount = 0;
				datax.forEach(function(elm){
					idx = arr.indexOf(elm);
					if (idx !== -1) {
						$(elements[idx]).parent().parent().remove();
						delContactCount += 1;
					}
				});
				// want to keep message for some time and then default
				notificationTextSet(delContactCount.toString()+
					" contacts Deleted Successfully...");
			},
			error: function(errorThrown){
				notificationTextSet(errorThrown);
			}
		});
		notificationTextSet();
	});

	$("#newContact form").submit(function(e){
		e.preventDefault();
		var formData = {
			'name':	 $("#name").val(),
			'contact': $("#contact").val(),
			'email_id': $("#email_id").val(),
			'address': $("#address").val()
		}
		$.ajax ({
			type: 'GET',
			data: formData,
			url: 'http://localhost:3000/insert',
			success: function(data) {
				notificationTextSet("new contact added.");
			}
		});
	});

	$("#filter").keyup(function() {
		var filter = $(this).val();
		var count = 0;
		var $rows = $('#contact-details table tr');
		var rowcount = $rows.length;

		$rows.each(function() {
			if ($(this).text().search(new RegExp(filter, "i")) < 0){
				$(this).fadeOut();
			}
			else{
				$(this).fadeIn();
				count++;
			}
			notificationTextSet(count);
		});

		if(rowcount == count || count == 0) {
			notificationTextSet();
		}
	});
});

$("#contact-details").submit(function(event){
    // this tells the server-side process that Ajax was used

    /* returns string which is url Query part*/
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
        dataType: 'javascript',
        success: function(data) {
            // want to keep message for some time and then default
            notificationTextSet();
        },
        error: function(errorThrown){
            notificationTextSet(errorThrown);
        }
    });
    /*
    var e = $('.del-selected').each(function(){
        var x = {m: [$(this).attr('recordno'), $(this)]};
        console.log(x);
    });
    */
});
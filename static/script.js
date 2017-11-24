//David Clarke G00335563
function scriptFunction() {
    const form = $("#inputUser");
    const listItem = $("#list-group");

    //stops the page from refreshing
    event.preventDefault();

    const input = form.val();

    form.val(" ");
    console.log("first")
    console.log(input)
    listItem.append("<li class='list-group-item list-group-item-primary text-left'>" + "User : " + input + "</li>");
    console.log("after")

    const queryParams = {
        "inputUser": input
    }
    $.get("/chat", queryParams)
        .done(function(resp) {
            //setting output to the right and making it red
            const nextListItem = "<li class='list-group-item list-group-item-warning text-right'>" + "Eliza : " + resp + "</li>";
            //        setTimeout(function() {
            listItem.append(nextListItem)
            //        $("html, body").scrollTop($("body").height());
            //  }, 1000); //set response wait time
        }).fail(function() { //Whenever the server is not connected, output the knock knock
            const nextListItem = "<li class='list-group-item list-group-item-danger text-right' >Eliza : Knock knock! Whos there? Not me!</li>";
            listItem.append(nextListItem);
        });


} //end of scriptFunction
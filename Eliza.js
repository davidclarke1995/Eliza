function myFunction(){
const form = $("#userInput");
const listItem = $("#list-group");

    console.log("Welcome to Eliza chat");
    event.preventDefault();

    const input = form.val();

    form.val(" ");
    console.log("before")
    console.log(input)
    listItem.append("<li class='list-group-item list-group-item-primary text-right'>" + "User : " + input + "</li>");
      
console.log("after")

        const queryParams = {
            "userInput": input
        }
        $.get("/chat", queryParams)
            .done(function(resp) {
                var nextListItem = "<li class='list-group-item list-group-item-info'>" + "ELiza : " + output + "</li>";
                setTimeout(function() {
                    listItem.append(nextListItem)
                }, 1000); //
            }).fail(function() {
                var nextListItem = "<li class='list-group-item list-group-item-danger' >Knock knock, who's there? Not me!</li>";
                listItem.append(nextListItem);
            });

}
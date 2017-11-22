console.log("before function")

const form = $("#user-input");
const list = $("#conversation_list");


form.keypress(function(event){
    if(event.keyCode != 13){ // ENTER
        return;
    }

    console.log("After function")

    event.preventDefault(); //Stops the page from refreshing

    const userText = form.val(); // get the text from the input form
    form.val(" "); // wipes the text box.
    
    // before you send request, make sure the user input is valid i.e. not all empty.
    list.append("<li  class='list-group-item  text-left list-group-item-success' id='leftList'>"+"User : " + userText + "</li>");

    // GET/POST
    const queryParams = {"inputUser" : userText }
    $.get("/chat", queryParams)
    
        .done(function(resp){
            //const newItem = "<img id='avatar-right' src='/img/eliza-avatar.png' alt='eliza-avatar'><li  class='list-group-item text-left list-group-item-info' id='rightList'>"+"ELiza : " + resp + "</li>";
            setTimeout(function(){
              
                list.append(newItem)
            }, 2000);//set timeout to give wait to response
        }).fail(function(){
            const newItem = "<li class='list-group-item list-group-item-danger' >Knock knock, who's there? Not me!</li>";
            list.append(newItem);
        });
});
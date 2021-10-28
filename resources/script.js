function CreateUser() {
    let xhr = new XMLHttpRequest();
    xhr.open("PUT", "/user");

    let input = document.querySelector("#UserName");
    let o = {
        name: input.value,
    };

    xhr.send(
        JSON.stringify(o)
    );
}
const back_button = document.getElementById("back_button");

back_button.addEventListener("click",function(evt){
    window.location.href = "./index.html"
})

async function main() {
    const test_input = document.getElementById('test_input');

    UseBootstrapTag(test_input);

    test_input.addEventListener("change",function(evt){
        console.log(test_input.value.split(","));
    })
}

main();
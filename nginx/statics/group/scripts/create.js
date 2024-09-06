const GetUserBtn = document.getElementById("GetUserBtn");

GetUserBtn.addEventListener("click",async function (evt) {
    console.log(await GetInfo());
})
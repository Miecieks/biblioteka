if(sessionStorage.getItem("id") != null){

    let id = sessionStorage.getItem("id")
    let avatar_inner = document.getElementById("avatar")
    let screener = document.getElementById("screenOfChange")
    let changePass = document.getElementById("changePass")

    let First_name = document.getElementById("FN")
    let Last_name = document.getElementById("LN")
    let email = document.getElementById("EM")
    let UiD = document.getElementById("UIDprof")
    
    let sub = document.getElementById("sub") 
    let name = document.getElementById("name") 
    let ev = document.getElementById("everything") 
    

    let closer = document.getElementById("closer")
    console.log(id)
    screener.style.display = "none"
    const infoUser = await GetUser(id)
    console.log(infoUser)
    let temp = "../images/"+infoUser.pfp
    avatar_inner.src = temp


    name.innerHTML = infoUser.user_name
    First_name.value = infoUser.first_name
    Last_name.value = infoUser.last_name
    email.value = infoUser.email
    UiD.value = infoUser.user_inner_id

    let logout_btn = document.getElementById("logout")
    logout_btn.addEventListener("click",logout)
    changePass.addEventListener("click",openScreen)
    closer.addEventListener("click",closeScreen)
    sub.addEventListener('click',change_password)
    function logout(){
        sessionStorage.clear()
        window.location.replace("../index.html")
    }

    function openScreen(){
        screener.style.display = "block"

    }

    function closeScreen() {
        screener.style.display = "none"
  
    }

    function change_password(){
        let newPass = document.getElementById("newPass")
        
        const passed = {
          Id: Number(id),
          Password: newPass.value
        };
        fetch("http://localhost:8080/api/user/passChange",{
            method: "POST",
            headers: {
              "Content-type": "application/json"
            },
            body: JSON.stringify(passed)
        })
        .then(res => res.json())
        .then(data => {
          if (data.success) {
            console.log("działa jak naturson chciał");
          }else{

            console.log(data)
          }
        })
    }
}else{

    window.location.replace("login.html")
    
}



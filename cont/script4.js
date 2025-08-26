if(sessionStorage.getItem("id") != null){
    let bar = document.getElementById("logNav")
    let avatar = document.getElementById("avNav")
    let avatar_inner = document.getElementById("avatar")
    bar.style.display = "none"
    avatar.style.display = "block"
    let id = sessionStorage.getItem("id")
    console.log(id)
    const infoUser = await GetUser(id)
    console.log(infoUser)
    let temp = "images/"+infoUser.pfp
    avatar_inner.src = temp
    let x = sessionStorage.getItem("adm")
    if(x == "true"){
        let navBarholder = document.getElementById("navBarholder")
        const li = document.createElement("li");
        const a = document.createElement("a");
        const TxNode = document.createTextNode("Zarządzanie Książkami");
        li.setAttribute("class","nav-item");
        a.setAttribute("class","nav-link");
        a.setAttribute("href","admin-side/book_managment.html")
        a.appendChild(TxNode);
        li.appendChild(a);
        navBarholder.appendChild(li);
    }
}else{
    let bar = document.getElementById("logNav")
    let avatar = document.getElementById("avNav")
    bar.style.display = "block"
    avatar.style.display = "none"
    
    
}



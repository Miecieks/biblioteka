if(sessionStorage.getItem("id") != null){
    let id = sessionStorage.getItem("id")
    console.log(id)
    const infoUser = await GetUser(id)
    console.log(infoUser)

}


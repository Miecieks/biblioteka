if(sessionStorage.getItem("adm") == "true"){
let main_cont = document.getElementById("main")
main_cont.style.display = "block";
let adder = document.getElementById("adder")
let selector = document.getElementById("Libraries")




fetch("http://localhost:8080/api/stuff/libraries/get/all")
  .then(res => res.json())
  .then(data => {
        if (data.success) {
            console.log(data.res[0])
            let i = 0
            while(data.res[i]){
                let option = document.createElement("option")
                option.setAttribute("value",data.res[i].Id)
                option.innerHTML = data.res[i].Name
                selector.appendChild(option)
                i += 1
            }
        }else{
            console.log("ciapa")
        }
    })
  .catch(err => console.log("Błąd połączenia:", err));



  let submit = document.getElementById("submit")
  

  submit.addEventListener("click",add)

  function add(){
      let book_name = document.getElementById("book_name").value
      let author = document.getElementById("author").value
      let price = document.getElementById("price").value
      let genre = document.getElementById("genre").value
      let library_id = document.getElementById("Libraries").value
      let acess = document.getElementById("acess").value
      let acess2 = (acess === 'true');
      if(book_name != ""){
      const book = {
          name: book_name,
          author: author,
          price: parseFloat(price),
          genre: genre,
          library_id: parseInt(library_id),
          is_avaible: acess2
        };
        console.log(book)
        fetch("http://localhost:8080/api/stuff/libraries/book/add", {
        method: "POST",
        headers: {
          "Content-type": "application/json"
        },
        body: JSON.stringify(book)
        })
        .then(res => res.json())
        .then(data => {
        console.log(data)
          console.log("Dodano książkę!");
          location.reload();
      })
    }

  }
}else{
    window.location.replace("../index.html")
}
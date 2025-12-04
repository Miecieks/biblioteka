let params = new URLSearchParams(document.location.search);
let id = params.get("id")
let rent = document.getElementById("renter")
let lib_id = 0
let book_id = 0
fetch("http://localhost:8080/api/stuff/libraries/book/get/"+id+"/extra")
  .then(res => res.json())
  .then(data => {
        if (data.success) {
            console.log(data.res)
            let name = document.getElementById("bookName")
            let author = document.getElementById("author")
            let genre = document.getElementById("genre")
            let library = document.getElementById("library")
            let cover = document.getElementById("cover")
            let avaib = document.getElementById("avaib")

            name.innerHTML = data.res.Name
            if(data.res.Is_avaible == "true"){
              avaib.innerHTML = "dostępna"
            }else{
              avaib.innerHTML = "niedostępna"
            }
            author.innerHTML = data.res.Author
            library.innerHTML = " "+data.res.LibName
            genre.innerHTML = " "+data.res.Genre
            cover.setAttribute("src","../"+data.res.Cover)

            book_id = data.res.Id
            lib_id = data.res.Library_id

        }else{
            console.log("ciapa")
        }
    })
  .catch(err => console.log("Błąd połączenia:", err));


  rent.addEventListener("click",renting)

  function renting(){

    


    let userId = sessionStorage.getItem("id") 
    console.log(book_id+"   "+userId)
    const rent = {
      Book_id: book_id,
      User_id: parseInt(userId),
    }
    fetch("http://localhost:8080/api/stuff/libraries/book/rent",{
      method: "POST",
      headers: {"Content-type": "application/json"},
      body: JSON.stringify(rent)
    })
    .then(data => {
      console.log(data)
      let texts = document.getElementById("texter")
      tango = data.status
      tango = parseInt(tango)
      if(tango == 321){
        texts.innerHTML = "Błąd! książka już jest wypożyczona!"
      }else if(data.status == "200"){
        texts.innerHTML = "Wypożyczono!"
      }

        })
  }
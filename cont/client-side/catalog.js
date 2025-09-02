let holder = document.getElementById("holder")






fetch("http://localhost:8080/api/stuff/libraries/book/get")
  .then(res => res.json())
  .then(data => {
        if (data.success) {
            console.log(data.res[0])
            let i = 0
            while(data.res[i]){
                let spandam = document.createElement("div");
                let img = document.createElement("img");
                let a = document.createElement("a");
                let p = document.createElement("p");
                img.setAttribute("id",data.res[i].Id);
                img.setAttribute("class","rounded float-start");
                spandam.setAttribute("class","float-start");
                img.setAttribute("src","../"+data.res[i].Cover);
                img.setAttribute("alt","ups! nie mamy takiego zdjęcia!");
                spandam.appendChild(img);
                holder.appendChild(spandam);
                
                a.innerHTML = data.res[i].Name
                a.setAttribute("href","product.html?id="+data.res[i].Id)

                p.appendChild(a)
                spandam.appendChild(p)
                i += 1
            }
        }else{
            console.log("ciapa")
        }
    })
  .catch(err => console.log("Błąd połączenia:", err));



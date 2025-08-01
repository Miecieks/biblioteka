async function getDataJson(url) {
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }

    const json = await response.json();
    let x = json
    return x
  } catch (error) {
    console.error("BŁĄD: " + error.message);
  }
}

async function GetUser(id) {
    
    let url = "http://localhost:8080/api/user/"+id+"/get";
    let x = await getDataJson(url)
    return x
    
}


function further(){
  if(document.getElementById("sub")){
    document.getElementById("first").style.display = "none"
    document.getElementById("second").style.display = "block"
  }
}

function further2(){
  if(document.getElementById("sub2")){
    document.getElementById("second").style.display = "none"
    document.getElementById("third").style.display = "block"
  }
}

function register(){
  console.log("rejestracja WIP")
}

function validate(){
  let user = document.getElementById("login").value
  let pass = document.getElementById("pass").value
  const LoginStruct = {
    login: user,
    pass: pass
  };

  fetch("http://localhost:8080/api/user/verify",{
    method: "POST",
    headers: {
      "Content-type": "application/json"
    },
    body: JSON.stringify(LoginStruct)
  })
  .then(res => res.json())
  .then(data => {
        if (data.success) {
      console.log("Zalogowano pomyślnie!");
      let info = document.getElementById("info")
      info.innerHTML = "Zalogowano pomyślnie! "
      sessionStorage.setItem("logged", true);
      sessionStorage.setItem("id", data.id);
    } else {
      console.log("Błąd logowania:", data.message);
      let info = document.getElementById("info")
      info.innerHTML = "Błąd logowania"
    }
  })
    .catch(err => console.error("Błąd połączenia:", err));
}

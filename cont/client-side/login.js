async function getDataJson(url) {
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }

    const json = await response.json();
    console.log(json);
    console.log(JSON.stringify(json))
    localStorage.setItem("username", url);
  } catch (error) {
    console.error("BŁĄD: " + error.message);
  }
}

function GetUser() {
    
    let user = document.getElementById("login").value
    let pass = document.getElementById("pass").value
    let url = "http://localhost:8080/api/user/"+user+"/get";
    let x = JSON.stringify(getDataJson(url))
    console.log(x)
    
    
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
      info.innerHTML = "Zalogowano pomyślnie!"
    } else {
      console.log("Błąd logowania:", data.message);
    }
  })
    .catch(err => console.error("Błąd połączenia:", err));
}

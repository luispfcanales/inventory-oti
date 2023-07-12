if (typeof idExitSession === "undefined") {


  function GET_COOKIE(name) {
    let cookies = document.cookie.split(';');
    for (let i = 0; i < cookies.length; i++) {
      let cookie = cookies[i].trim();
      if (cookie.startsWith(name + '=')) {
        return cookie.substring(name.length + 1);
      }
    }
    return null;
  }
  
  const idExitSession = document.getElementById("event-exit-session")
  idExitSession.addEventListener("click",()=>{
    //let token = GET_COOKIE("Authorization")
    //alert(atob(token.split('.')[1]))

    document.cookie = "Authorization=;"
    window.location.href = `${URL_APP}/login`
  })

}

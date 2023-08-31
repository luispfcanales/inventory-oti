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


const ExitSession = document.getElementById("btnExit")
ExitSession.addEventListener("click",()=>{
  window.location.href = "/exit"
})


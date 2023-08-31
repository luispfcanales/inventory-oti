let METHOD_HTTP_MODAL ="POST"

function CleanDataModal() {
  METHOD_HTTP_MODAL ="POST"

  title_modal.innerText = "Agregar nuevo usuario"
  input_dni.value=""
  firstName.value = ""
  lastName.value = ""
  input_email.value = ""
  select_role.options[0].selected = true
  select_state.options[0].selected = true
  select_staff.options[0].selected = true

  error_message.classList.add('hidden');
  document.querySelector('#error_message span').innerText = ""
}

function InputsModalDisable(value,disableDNI = false){
  input_dni.disabled = disableDNI
  select_state.disabled = value
  select_role.disabled = value
  select_staff.disabled = value
  input_email.disabled = value
  btn_modalUser.disabled = value
}

//methods update values person
function selectedItemToTrue(id,item){
  item.map(option => {
    if (id == option.value){
      option.selected = true
    }
  })
}
function SetValuesModalToUpdate({key,dni,email,first_name,last_name,id_role,id_staff,state}){
  METHOD_HTTP_MODAL = "PUT"
  input_dni.value = dni 
  title_modal.innerText = "Actualizar Datos"
  firstName.value = first_name
  lastName.value = last_name
  input_email.value = email

  selectedItemToTrue(id_role,[...select_role.options])
  selectedItemToTrue(id_staff,[...select_staff.options])
  selectedItemToTrue(state,[...select_state.options])

  InputsModalDisable(false,true)
}


//methods get values person by DNI API
function SetValuesApiPerson(data){
  firstName.value = data.nombres
  lastName.value = `${data.apellidoPaterno} ${data.apellidoMaterno}`
}
const GetPersonByDNI= async(dni) => {
  let url = `${URL_API}/dni/${dni}`
  document.querySelector("#loaderModal").classList.toggle('hidden')
  try{
    const request = await fetch(url)
    const data = await request.json()

    console.log(data)
    if(data.status === 404 ){
      error_message.classList.remove('hidden');
      document.querySelector('#error_message span').innerText = data.data.message
      document.querySelector("#loaderModal").classList.toggle('hidden')
      return
    }

    SetValuesApiPerson(data.data)
    InputsModalDisable(false)
    error_message.classList.add('hidden');
    document.querySelector("#loaderModal").classList.toggle('hidden')
  }catch(err){
    console.log(err)
    document.querySelector("#loaderModal").classList.toggle('hidden')
  }
}
const inputDni = document.getElementById("input_dni")
inputDni.addEventListener("input",()=>{
  inputDni.value = inputDni.value.replace(/[^0-9]/g, "");
  if(inputDni.value.length < 8){
    InputsModalDisable(true)
    firstName.value = ""
    lastName.value = ""
    return
  }
  GetPersonByDNI(inputDni.value)
})


const userRegister=async(user)=>{
  let url = `${URL_API}/users`
  document.querySelector("#loaderModal").classList.toggle('hidden')

  try{
    const request = await fetch(url,{
      method: METHOD_HTTP_MODAL,
      headers:{
        'Content-Type':'application/json'
      },
      body: JSON.stringify(user)
    })
    const body = await request.json()

    if(body.status == 400) {
      if(typeof body.data != undefined){
        error_message.classList.remove('hidden');
        document.querySelector('#error_message span').innerText = body.data
      }
      document.querySelector("#loaderModal").classList.toggle('hidden')
      return
    }

    error_message.classList.add('hidden');
    document.querySelector("#loaderModal").classList.toggle('hidden')
    document.querySelector("#container_modal_user").classList.toggle('hidden')
    CleanDataModal()
    InputsModalDisable(true)

  }catch(err){
    console.log(err)
    document.querySelector("#loaderModal").classList.toggle('hidden')
  }
}

btn_modalUser.addEventListener("click",()=>{
  const user = {
    dni:input_dni.value,
    first_name:firstName.value,
    last_name:lastName.value,
    id_role:select_role.value,
    id_staff:select_staff.value,
    password:"12345678",
    state:parseInt(select_state.value),
    email:input_email.value,
  }
  userRegister(user)

})

const open_modal_user = document.getElementById("open_modal_user")
open_modal_user.addEventListener("click",()=>{
  METHOD_HTTP_MODAL ="POST"
  document.querySelector("#container_modal_user").classList.toggle('hidden')
})
const close_modal_user = document.getElementById("close_modal_user")
close_modal_user.addEventListener("click",()=>{
  document.querySelector("#container_modal_user").classList.toggle('hidden')
  CleanDataModal()
  InputsModalDisable(true)
})

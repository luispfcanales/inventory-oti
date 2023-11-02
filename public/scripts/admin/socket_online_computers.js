let socket = new WebSocket(`${STREAM_COMPUTERS}/3gr3df`)
let hashmap = new Map()

const EVENT_NOTIFY = 1//"notify"
const EVENT_LOAD_INFO = 2//"load-info-system"
const EVENT_LOADED = 3//"loaded-info"

const loadOnlineDesktops = async() => {
  const url = `${STREAM_URL}`
  const request = await fetch(url)
  const body = await request.json()

  if(body.data === null){
    return
  }
  body.data.map(data=>{
    hashmap.set(data.id,data)
  })
  renderValues()
}

function renderValues() {
  let info = []
  let nro = 0
  hashmap.forEach(function(value){
    nro = nro + 1;
    info = [...info,`
        <tr class="hover:bg-gray-100">
          <td class="text-center border-b">${nro}</td>
          <td class="text-center border-b">${value.id}</td>
          <td class="text-center border-b">${value.status}</td>
          <td class="text-center px-3 text-center border-b">
            <div class="flex justify-center gap-4 group-hover:block w-full" role="menu" aria-orientation="vertical" aria-labelledby="menu-button" tabindex="-1">
                <button
                  onclick="LoadModalWithID('${value.id}')"
                  class="flex text-cyan-700 block text-sm" role="menuitem" tabindex="-1" id="menu-item-2"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M9.348 14.651a3.75 3.75 0 010-5.303m5.304 0a3.75 3.75 0 010 5.303m-7.425 2.122a6.75 6.75 0 010-9.546m9.546 0a6.75 6.75 0 010 9.546M5.106 18.894c-3.808-3.808-3.808-9.98 0-13.789m13.788 0c3.808 3.808 3.808 9.981 0 13.79M12 12h.008v.007H12V12zm.375 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z" />
                  </svg>
                </button>
                <button
                  onclick="alert('delete')"
                  class="flex hover:bg-gray-100 text-red-700 block text-sm" role="menuitem" tabindex="-1" id="menu-item-3">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M11.42 15.17L17.25 21A2.652 2.652 0 0021 17.25l-5.877-5.877M11.42 15.17l2.496-3.03c.317-.384.74-.626 1.208-.766M11.42 15.17l-4.655 5.653a2.548 2.548 0 11-3.586-3.586l6.837-5.63m5.108-.233c.55-.164 1.163-.188 1.743-.14a4.5 4.5 0 004.486-6.336l-3.276 3.277a3.004 3.004 0 01-2.25-2.25l3.276-3.276a4.5 4.5 0 00-6.336 4.486c.091 1.076-.071 2.264-.904 2.95l-.102.085m-1.745 1.437L5.909 7.5H4.5L2.25 3.75l1.5-1.5L7.5 4.5v1.409l4.26 4.26m-1.745 1.437l1.745-1.437m6.615 8.206L15.75 15.75M4.867 19.125h.008v.008h-.008v-.008z" />
                  </svg>
                </button>
            </div>
          </td>
        </tr>
  `]
  })

  $("#body_table_online_computer").html(info)
}

loadOnlineDesktops()

const LoadModalWithID =(PID)=>{
  const data = { id: PID, status: "online", event:EVENT_LOAD_INFO , role:"admin", event_emisor_id:"3gr3df" };
  const json = JSON.stringify(data);
  socket.send(json)
  //document.querySelector("#container_modal_online_computers").classList.toggle('hidden')
}
function processNotifyEvent(obj){
  switch(obj.status){
    case "online":
      hashmap.set(obj.id,obj)
      break
    case "offline":
      hashmap.delete(obj.id)
      break
  }
  renderValues()
}

const processEvent =(obj)=>{
  switch(obj.event){
    case EVENT_LOADED:
      console.log(obj)
      break
    case EVENT_NOTIFY:
      processNotifyEvent(obj)
      break
  }
}

socket.addEventListener("message",(e)=>{
  let obj = JSON.parse(e.data)
  processEvent(obj)
})


const send = document.getElementById("testsocket")
send.addEventListener("click",()=>{
})

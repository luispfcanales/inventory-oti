const fxHoja = document.getElementById("fxhoja")

let { pdfjsLib } = globalThis;
//let pdfjsLib = window['pdfjs-dist/build/pdf'];

pdfjsLib.GlobalWorkerOptions.workerSrc = '//mozilla.github.io/pdf.js/build/pdf.worker.mjs';
//pdfjsLib.GlobalWorkerOptions.workerSrc
//https://mozilla.github.io/pdf.js/build/pdf.worker.mjs

let loadingtask = pdfjsLib.getDocument("/preview/pdf")

const factor = 2.5
const datos_firma = {
  archivo_id:0,
  num_pagina:0,
  motivo:'soy el autor',
  exacto:1,
  pos_pagina:"0-0",
  apariencia:0
}

function printpdf() {
  //const urlblob = URL.createObjectURL(vblob)
  //const url = `${urlblob}.pdf`

  loadingtask.promise.then(function(pdf){
    pdf.getPage(1).then(function(page){
      let viewport = page.getViewport({scale:.75})
      //let viewport = page.getViewport({scale:1})
      let canvas = document.createElement('canvas')
      let ctx = canvas.getContext('2d')
      //canvas.width = 595;
      //canvas.height = 842;
      canvas.width = 446.25;
      canvas.height = 631.5;

      page.render({canvasContext:ctx,viewport:viewport}).promise.then(function(){
        const imgcanvas = canvas.toDataURL('image/png')
        fxHoja.style.position = "relative";
        fxHoja.style.width = canvas.width+"px";
        fxHoja.style.height = canvas.height+"px";
        fxHoja.style.backgroundImage = `url("${imgcanvas}")`
      })
    })
  })

}

async function LoadPdf() {
  //const r = await fetch(url)//.then(response => response.atob())
  //const value = await r.blob()
  printpdf()
}
LoadPdf()



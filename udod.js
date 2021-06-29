var state = {
    item: [
        {
            text: '',done : false
        }
    ]

}
window.addEventListener('load',(Event) => {
    fetch("http://localhost:8080/getAll")
    .then(res => res.json())
    .then(json => {
        state = json
        console.log(state);
        drawAll()
    })

})


function newItem() {
    text = item.text
    done = item.done
    document.getElementById("2").innerHTML += `
    <div class = "red"><div class = "white">
    ${document.getElementById("1").value}
    </div> 
    <div class = "white">
        <input type="checkbox">
    </div>
    </div>`
}

function drawItem(item) {
    text = item.text
    done = item.done
    document.getElementById("2").innerHTML += `
    <div class = "red"><div class = "white">
    ${text}
    </div> 
    <div class = "white">
        <input type="checkbox" checked = "${done}">
    </div>
    </div>`
}

function drawAll() {
    for(i = 0;i < state.notes.length; i++){
        drawItem(state.notes[i])
        

    }
}


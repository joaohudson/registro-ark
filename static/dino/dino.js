const h1 = document.getElementById('h1');
const title = document.getElementById('title');
const regionP = document.getElementById('regionP');
const locomotionP = document.getElementById('locomotionP');
const foodP = document.getElementById('foodP');
const utilityP = document.getElementById('utilityP');
const trainingP = document.getElementById('trainingP');

async function getDino(id){
    const response = await fetch('/api/dino?id=' + id);

    if(!response.ok){
        throw await response.text();
    }

    return await response.json();
}

async function main(){
    const params = new URLSearchParams(location.search);
    if(!params.has('id')){
        alert('Dino não encontrado!');
        return;
    }
    const dinoId = params.get('id');

    
    let dino;
    try{
        dino = await getDino(dinoId)
    }
    catch(e){
        alert(e);
        return;
    }

    h1.innerText = dino.name;
    title.innerText = 'Registro Ark - ' + dino.name;
    regionP.innerText = dino.region;
    locomotionP.innerText = dino.locomotion;
    foodP.innerText = dino.food;
    utilityP.innerText = dino.utility;
    trainingP.innerText = dino.training;
}

main();
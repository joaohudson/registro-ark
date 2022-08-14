const h1 = document.getElementById('h1');
const title = document.getElementById('title');
const infoDiv = document.getElementById('infoDiv');
const regionP = document.getElementById('regionP');
const locomotionP = document.getElementById('locomotionP');
const foodP = document.getElementById('foodP');
const utilityP = document.getElementById('utilityP');
const trainingP = document.getElementById('trainingP');

//image
const dinoImage = document.getElementById('dinoImage');
const dinoImageDiv = document.getElementById('dinoImageDiv');

async function getDino(id){
    const response = await fetch('/api/dino?id=' + id);

    if(!response.ok){
        throw await response.text();
    }

    return await response.json();
}

async function getImageByDino(id){
    const response = await fetch('/api/dino/image?id=' + id);

    if(!response.ok){
        throw await response.text();
    }

    return await response.json();
}

function showError(message){
    const p = document.createElement('p');
    p.innerText = message;
    infoDiv.innerHTML = '';
    infoDiv.appendChild(p);
}

async function main(){
    const params = new URLSearchParams(location.search);
    if(!params.has('id')){
        showError('Dino não encontrado!');
        return;
    }
    const dinoId = params.get('id');

    
    let dino;
    try{
        dino = await getDino(dinoId)
    }
    catch(e){
        showError(e);
        return;
    }
    
    h1.innerText = dino.name;
    title.innerText = 'Registro Ark - ' + dino.name;
    regionP.innerText = dino.region.name;
    locomotionP.innerText = dino.locomotion.name;
    foodP.innerText = dino.food.name;
    utilityP.innerText = dino.utility;
    trainingP.innerText = dino.training;
    
    try{
        const image = await getImageByDino(dinoId);
        dinoImageDiv.hidden = false;
        dinoImage.src = image.base64;
    }catch(_){}
}

main();
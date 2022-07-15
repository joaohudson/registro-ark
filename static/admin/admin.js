(function (){

const nameLabel = document.getElementById('nameLabel');

const addDinoRoute = document.getElementById('addDinoRoute');
const listDinoRoute = document.getElementById('listDinoRoute');
const managerCategoryRoute = document.getElementById('managerCategoryRoute');
const managerAdmRoute = document.getElementById('managerAdmRoute');
const logoutButton = document.getElementById('logoutButton');

async function fetchAdm(){
    return await adminFetch('/api/adm', {}, true);
}

logoutButton.onclick = () => {
    delete localStorage.token;
    location.href = '/admin/login';
};

async function main(){
    try{
        const adm = await fetchAdm();

        nameLabel.innerText = 'Adm: ' + adm.name;

        addDinoRoute.hidden = !adm.permissionManagerDino;
        listDinoRoute.hidden = !adm.permissionManagerDino;

        managerCategoryRoute.hidden = !adm.permissionManagerCategory;

        managerAdmRoute.hidden = !adm.permissionManagerAdm;
    }
    catch(e){
        dialog.showMessage(e);
        return;
    }
}

main();

})();
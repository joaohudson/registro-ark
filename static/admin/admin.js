(function (){

const token = localStorage.getItem('token');
if(!token){
    location.href = '/admin/login';
}

})();
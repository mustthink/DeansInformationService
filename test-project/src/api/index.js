import { LoginAPIInstance, DefaultAPIInstance } from "./AuthAPI";

export const AuthAPI = {

    /** 
    * @param {string} login
    * @param {string} password
    * @returns {Promise<AxiosResponse<any>>}
    */

    login(login, password){
        const url = '/login';
        const data = {login, password};
        return LoginAPIInstance(url, data);
    },

    /**
     * 
     * @returns {Promise<AxiosResponse<any>>}
     */
    
    logout(){
        const url = '/logout';
        return DefaultAPIInstance.post(url);
    }
}
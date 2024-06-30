import { useState } from "react";
import Dropdown from "../components/dropDown";
import styles from "./styles/index";
import { ConnectToDB } from "../../wailsjs/go/main/App"
import { useNavigate } from "react-router-dom";

function Index() {
    const [host, setHost] = useState("");
    const [port, setPort] = useState("");
    const [user, setUser] = useState("");
    const [password, setPassword] = useState("");
    const [dbname, setDbname] = useState("");
    const [selectedOption, setSelectedOption] = useState<number>(-1);
    const navigate = useNavigate();

    const handleConnect = async () => {
        // ConnectToDB(host, port, user, password, dbname, selectedOption).then((_) => {
        //     navigate("/insert");
        // }).catch((_) => {
        //     console.log("Error");
        // });
        navigate("/insert");
    }

    return (
        <div className="grid grid-rows-5 grid-cols-5 w-full h-full bg-cyan-400">
            <h1 className={styles.home}>Home</h1>
            <div className="grid grid-rows-5 row-start-2 row-span-3 col-start-2 col-span-3">
                <div className="grid grid-cols-3 row-start-1 row-end-2">
                    <label htmlFor="host-bar" className={styles.label}>Host :</label>
                    <input type="text" id="host-bar" className={styles.input} 
                        value={host} onChange={(e) => setHost(e.target.value)}/>
                </div>
                <div className="grid grid-cols-3 row-start-2 row-end-3">
                    <label htmlFor="port-bar" className={styles.label}>Port :</label>
                    <input type="text" id="port-bar" className={styles.input} 
                        value={port} onChange={(e) => setPort(e.target.value)}/>
                </div>
                <div className="grid grid-cols-3 row-start-3 row-end-4">
                    <label htmlFor="user-bar" className={styles.label}>User :</label>
                    <input type="text" id="user-bar" className={styles.input} 
                        value={user} onChange={(e) => setUser(e.target.value)}/>
                </div>
                <div className="grid grid-cols-3 row-start-4 row-end-5">
                    <label htmlFor="password-bar" className={styles.label}>Password :</label>
                    <input type="password" id="password-bar" className={styles.input}
                         value={password} onChange={(e) => setPassword(e.target.value)}/>
                </div>
                <div className="grid grid-cols-3 row-start-5 row-end-6">
                    <label htmlFor="dbname-bar" className={styles.label}>DB Name :</label>
                    <input type="text" id="dbname-bar" className={styles.input}
                         value={dbname} onChange={(e) => setDbname(e.target.value)}/>
                </div>
                <p className="text-xs font-semibold italic font-mono ml-14 w-full">
                    ( For sqlite, it is the path; For mongodb, it is in format [tablename/collection name] )
                </p>
            </div>
            <div className="grid grid-cols-2 row-start-5 row-end-6 col-start-2 col-span-3 m">
                <Dropdown position="col-start-1 col-end-2 ml-16" selectedOption={selectedOption} setSelectedOption={setSelectedOption}/>
                <button className={styles.button} onClick={handleConnect}>Connect</button>
            </div>
        </div>
    );
}

export default Index;
import { faMinus, faPlus } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useState } from "react";

interface pair {
    key: string;
    value: string;

}

function InsertPage() {
    const [elements, setElements] = useState<pair[]>([{ key: '', value: '' }]);

    const handleAddElement = () => {
        setElements([...elements, { key: '', value: '' }]);
    };

    const handleRemoveElement = (index: number) => {
        const newElements = elements.filter((_, i) => i !== index);
        setElements(newElements);
    };

    const handleChange = (index: number, event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        const newElements = elements.map((element, i) => {
            if (i === index) {
                return { ...element, [name]: value };
            }
            return element;
        });
        setElements(newElements);
    };

    const handleSubmit = (event: { preventDefault: () => void; }) => {
        event.preventDefault();
        console.log('Submitted Elements:', elements);
    };

    return (
        <div className="grid grid-rows-10 w-full h-full bg-cyan-400 overflow-hidden">
            <h1 className="flex justify-center items-center row-start-1 row-end-2 text-4xl font-bold font-mono">Insert Page</h1>
            <div className="row-start-2 row-span-7 my-1 w-full h-full overflow-y-auto">
                {elements.map((element, index) => (
                    <div key={index} className="ml-32 my-4">
                        <input
                            type="text"
                            name="key"
                            value={element.key}
                            className="w-60 px-2 py-1 mx-2 rounded-md"
                            placeholder="key"
                            onChange={(e) => handleChange(index, e)}
                        />: 
                        <input
                            type="text"
                            name="value"
                            value={element.value}
                            className="w-60 px-2 py-1 mx-2 rounded-md"
                            placeholder="value"
                            onChange={(e) => handleChange(index, e)}
                        />
                        <button type="button" className="ml-4 hover:text-red-600" onClick={() => handleRemoveElement(index)}>
                            <FontAwesomeIcon icon={faMinus} size="xl" />
                        </button>
                    </div>
                ))}
            </div>
            <div className="row-start-9 row-span-1 ml-80 mt-4">
                <button onClick={handleSubmit} 
                    className="text-3xl text-gray-800 -mx-6 mr-1 bg-green-400 border-2 border-solid border-black rounded-md py-2 px-4">
                    Submit
                </button>
                <button type="button" className="ml-10 hover:text-slate-600" onClick={handleAddElement}>
                    <FontAwesomeIcon icon={faPlus} size="2x" />
                </button>
            </div>
        </div>
    );
}

export default InsertPage;
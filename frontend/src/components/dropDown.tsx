import { useState } from 'react';
import styles from './styles/dropDown';

interface DropdownProps {
    position: string;
    selectedOption: number;
    setSelectedOption: (option: number) => void;
}

const Dropdown = ({ position, selectedOption, setSelectedOption }: DropdownProps) => {
    const options = ["PostgreSQL", "MySQL", "SQLite", "MongoDB"];
    const buttonClassName = `${styles.button} ${position}`;
    const [isOpen, setIsOpen] = useState<boolean>(false);

    const toggleDropdown = () => setIsOpen(!isOpen);

    const handleOptionClick = (option: number) => {
        setSelectedOption(option);
        setIsOpen(false);
    };

    return (
        <div className="relative inline-block text-left w-full flex justify-center items-center">
        <div className="w-56">
            <button
                type="button"
                className={buttonClassName}
                id="options-menu"
                aria-expanded="true"
                aria-haspopup="true"
                onClick={toggleDropdown}
            >
            {options[selectedOption] || "Select an Database Type"}
            <svg
                className="-mr-1 ml-2 h-5 w-5 rotate-180"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
                fill="currentColor"
                aria-hidden="true"
            >
                <path
                fillRule="evenodd"
                d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
                clipRule="evenodd"
                />
            </svg>
            </button>
        </div>

        {isOpen && (
            <div
            className={styles.list}
            >
                <div
                    className="py-1"
                    role="menu"
                    aria-orientation="vertical"
                    aria-labelledby="options-menu"
                >
                    {options.map((option, index) => (
                    <button
                        key={index}
                        onClick={() => handleOptionClick(index)}
                        className={styles.option}
                        role="menuitem"
                    >
                        {option}
                    </button>
                    ))}
                </div>
            </div>
        )}
        </div>
    );
};

export default Dropdown;

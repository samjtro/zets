import { useState } from "react"

const Display = ({counter}) => <div>{counter}</div>

const Button = ({onClick, text}) => <button onClick={onClick}>{text}</button>

/*1c
const App = () => {
    const [counter, setCounter] = useState(0)
    const increaseCounterByOne = () => setCounter(counter + 1)
    const decreaseCounterByOne = () => setCounter(counter - 1)
    const resetCounter = () => setCounter(0)
    return (
        <>
            <Display counter={counter} />
            <Button onClick={decreaseCounterByOne} text="-1" />
            <Button onClick={resetCounter} text="0" />
            <Button onClick={increaseCounterByOne} text="+1"/>
        </>
    )
}*/

const History = (props) => {
    if (props.allClicks.length === 0) {
        return (
            <div>
                press a button to continue
            </div>
        )
    }
    return (
        <div>
            history: {props.allClicks.join(', ')}
        </div>
    )
}

const App = () => {
    const [left, setLeft] = useState(0)
    const [right, setRight] = useState(0)
    const [allClicks, setAll] = useState([])
    const [total, setTotal] = useState(0)
    const handleLeftClick = () => {
        setAll(allClicks.concat('L'))
        const newLeft = left + 1
        setLeft(newLeft)
        setTotal(newLeft + right)
    }
    const handleRightClick = () => {
        setAll(allClicks.concat('R'))
        const newRight = right + 1
        setRight(newRight)
        setTotal(left + newRight)
    }
    return (
        <>
            <Display counter={left} />
            <Button onClick={handleLeftClick} text="left+1" />
            <Display counter={right} />
            <Button onClick={handleRightClick} text="right+1" />
            <Display counter={total} />
            <History allClicks={allClicks} />
        </>
    )
}

export default App
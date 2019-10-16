import React from 'react'

interface Props {
  currentTime: number
  duration: number
}

const ProgressBar: React.SFC<Props> = ({ currentTime, duration }) => {
  return (
    <div className="relative w-full h-1 bg-gray-400 rounded-full">
      <div
        className="absolute top-0 left-0 h-1 bg-red-500 rounded-full"
        style={{ width: `${(currentTime / duration) * 100}%` }}
      />
    </div>
  )
}

export default ProgressBar

import React from 'react'

interface Props {
  loadMore: () => void
  isLoading: boolean
}

const ButtonShowMore: React.SFC<Props> = (props) => {
  const { loadMore, isLoading } = props

  return (
    <button
      className="w-full h-full text-gray-800 bg-blue-200 rounded-full focus:outline-none focus:shadow-outline"
      onClick={loadMore}
    >
      {isLoading ? <div className="spinner mx-auto" /> : 'SHOW MORE'}
    </button>
  )
}

export default ButtonShowMore

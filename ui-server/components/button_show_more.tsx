import React from 'react'

interface Props {
  loadMore: () => void
  isLoading: boolean
}

const ButtonShowMore: React.SFC<Props> = (props) => {
  const { loadMore, isLoading } = props

  return (
    <button
      className="w-full h-full text-sm text-gray-900 tracking-tight bg-gray-200 rounded-lg focus:outline-none"
      onClick={loadMore}
    >
      {isLoading ? <div className="spinner mx-auto" /> : 'SHOW MORE'}
    </button>
  )
}

export default ButtonShowMore

import classNames from 'classnames'
import React from 'react'

export interface StateToProps {
  searchQuery: string
}

export interface OwnProps {
  resultType: 'podcast' | 'episode'
  sortBy: 'relevance' | 'publish_date'
}

const SearchResultsFilter: React.FC<StateToProps & OwnProps> = ({
  resultType,
}) => {
  return (
    <div className="flex mb-8">
      <div className="flex flex-initial w-1/2 border-b">
        <div className="w-20 mr-2 text-center">
          <div
            className={classNames(
              'block px-3 py-1 text-sm capitalize leading-loose tracking-wider',
              {
                'cursor-default': resultType === 'episode',
                'cursor-pointer': resultType !== 'episode',
              },
            )}
          >
            {'episode'}
          </div>
          <div
            className={classNames('h-1 w-20 rounded-full', {
              'bg-yellow-600': resultType === 'episode',
            })}
          />
        </div>
        <div className="w-20 mr-2 text-center">
          <div
            className={classNames(
              'block px-3 py-1 text-sm capitalize leading-loose tracking-wider',
              {
                'cursor-default': resultType === 'podcast',
                'cursor-pointer': resultType !== 'podcast',
              },
            )}
          >
            {'podcast'}
          </div>
          <div
            className={classNames('h-1 w-20 rounded-full', {
              'bg-yellow-600': resultType === 'podcast',
            })}
          />
        </div>
      </div>
    </div>
  )
}

export default SearchResultsFilter

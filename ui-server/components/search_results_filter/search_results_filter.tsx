import classNames from 'classnames'
import React from 'react'
import { SearchResultType, SearchSortBy } from 'types/search'

export interface StateToProps {
  searchBarText: string
}

export interface OwnProps {
  resultType: SearchResultType
  sortBy: SearchSortBy
}

const resultTypes: SearchResultType[] = ['episode', 'podcast']

const SearchResultsFilter: React.FC<StateToProps & OwnProps> = (props) => {
  return (
    <div className="flex mb-8">
      <div className="flex flex-initial w-1/2 border-b">
        {resultTypes.map((t) => (
          <div className="w-20 mr-2 text-center">
            <div
              className={classNames(
                'block px-3 py-1 text-sm capitalize leading-loose tracking-wider',
                {
                  'cursor-default': props.resultType === t,
                  'cursor-pointer': props.resultType !== t,
                },
              )}
            >
              {t}
            </div>
            <div
              className={classNames('h-1 w-20 rounded-full', {
                'bg-yellow-600': props.resultType === t,
              })}
            />
          </div>
        ))}
      </div>
    </div>
  )
}

export default SearchResultsFilter

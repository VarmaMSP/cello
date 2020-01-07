import classNames from 'classnames'
import React from 'react'
import { SearchResultType, SearchSortBy } from 'types/search'

export interface StateToProps {
  resultType: SearchResultType
  sortBy: SearchSortBy
}

export interface DispatchToProps {
  setResultType: (t: SearchResultType) => void
  setSortBy: (s: SearchSortBy) => void
  loadResultsPage: () => void
}

const SearchResultsFilter: React.FC<StateToProps & DispatchToProps> = (
  props,
) => {
  const onResultTypeChange = (t: SearchResultType) => {
    if (props.resultType !== t) {
      props.setResultType(t)
      props.loadResultsPage()
    }
  }

  const onSortByChange = (s: SearchSortBy) => {
    if (props.sortBy !== s) {
      props.setSortBy(s)
      props.loadResultsPage()
    }
  }

  return (
    <div className="flex items-center mb-8 justify-between">
      <div className="flex flex-initial w-3/5 border-b">
        {(['episode', 'podcast'] as SearchResultType[]).map((t) => (
          <div key={t} className="w-20 mr-2 text-center">
            <div
              className={classNames(
                'block px-3 py-1 text-sm capitalize leading-loose tracking-wider',
                {
                  'cursor-default': props.resultType === t,
                  'cursor-pointer': props.resultType !== t,
                },
              )}
              onClick={() => onResultTypeChange(t)}
            >
              {`${t}s`}
            </div>
            <div
              className={classNames('h-1 w-20 rounded-full', {
                'bg-green-600': props.resultType === t,
              })}
            />
          </div>
        ))}
      </div>
      <div>
        <select
          className="form-select text-sm tracking-wider w-36"
          onChange={(e) => onSortByChange(e.target.value as SearchSortBy)}
        >
          <option value="relevance">{'Relevance'}</option>
          {props.resultType === 'episode' && (
            <option value="publish_date">{'Date Published'}</option>
          )}
        </select>
      </div>
    </div>
  )
}

export default SearchResultsFilter

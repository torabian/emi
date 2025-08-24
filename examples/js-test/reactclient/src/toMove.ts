// export const getSinglePostUseQueryOptions = (
//   options: GetSinglePostOptions
// ) => ({
//   queryKey: [FetchGetSinglePostAction.NewUrl(options.params, options.qs)],
//   queryFn: () =>
//     FetchGetSinglePostAction.Fetch(options.params, options.qs, {
//       headers: options.headers,
//     }),
// });

// // 2️⃣ Hook: uses the query options with useQuery
// export const useGetSinglePost = (options: GetSinglePostQueryOptions) => {
//   return useQuery({
//     ...options,
//     ...getSinglePostQueryOptions(options),
//   });
// };

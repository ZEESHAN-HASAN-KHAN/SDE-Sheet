class Solution
{
public:
    int searchInsert(vector<int> &nums, int target)
    {
        // int ans=0;

        // for(int i=0;i<nums.size();i++)
        // {
        //     if(nums[i]==target)
        //     {
        //         ans=i;
        //         break;
        //     }
        //     else if(nums[i]<target)
        //     {
        //         ans=i+1;
        //     }
        //     else
        //     {
        //         break;
        //     }
        // }
        // return ans;
        int n = nums.size();
        int ans;
        int low = 0, high = n - 1;

        while (low <= high)
        {
            int mid = low + (high - low) / 2;

            if (nums[mid] == target)
            {
                ans = mid;
                break;
            }
            else if (nums[mid] < target)
            {
                ans = mid + 1;
                low = mid + 1;
            }
            else
            {
                high = mid - 1;
            }
        }
        return ans;
    }
};
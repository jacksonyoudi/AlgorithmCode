from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def helper(nodes: List[TreeNode], result: List[List[int]]) -> List[int]:
    new_nodes = []
    tmp = []
    for node in nodes:
        tmp.append(node.val)
        if node.left:
            new_nodes.append(node.left)
        if node.right:
            new_nodes.append(node.right)
    if len(new_nodes) == 0:
        return tmp
    else:
        tmp = helper(new_nodes, result)
        result.append(tmp)


class Solution:
    def levelOrderBottom(self, root: TreeNode) -> List[List[int]]:
        result = []
        helper([root], result)
        return result



select
concat((cast(regexp_replace(dt,'-','') as bigint) - 20190000) * 10000000 + row_num) AS id,
analyze_time,
shop_id,
total_consult_buyer,
total_consult_presale_buyer,
total_consult_aftersale_buyer,
total_send_rating_card,
total_evaluated_rating_card,
unix_timestamp() as ctime,
unix_timestamp() as mtime
from
(
    select *, row_number() over (partition by dt order by dt,analyze_time,shop_id) as row_num
from dw_shopee.aggr_webchat_shop_data_v2_di where dt='2020-06-03' limit 100
) as shop_data


SELECT
concat((cast(regexp_replace(dt,'-','') as bigint) - 20190000) * 10000000 + row_num)  AS id,
analyze_time,
shop_id,
subaccount_id,
average_serve_time,
total_consult_buyer,
average_waiting_time,
average_answer_time,
total_send_rating_card,
total_expire_rating_card,
total_evaluated_rating_card,
average_evaluated_rating_card_score,
grade,
reason,
unix_timestamp() as ctime,
unix_timestamp() as mtime
FROM (
    SELECT *, row_number() over (partition by dt order by dt,analyze_time,shop_id,subaccount_id desc) as row_num
from dw_shopee.aggr_webchat_subaccount_data_v2_di where dt='2020-06-03' and  average_evaluated_rating_card_score > 0 limit 100
) as subaccount_data
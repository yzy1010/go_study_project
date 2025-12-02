START TRANSACTION;

-- 获取账户A的余额并加锁
SET @balance_a = (SELECT balance FROM accounts WHERE id = 1 FOR UPDATE);

-- 检查余额是否足够
IF @balance_a >= 100 THEN
    -- 扣除账户A的金额
    UPDATE accounts SET balance = balance - 100 WHERE id = 1;

    -- 增加账户B的金额
    UPDATE accounts SET balance = balance + 100 WHERE id = 2;

    -- 记录交易
    INSERT INTO transactions (from_account_id, to_account_id, amount)
    VALUES (1, 2, 100);

    -- 提交事务
    COMMIT;
ELSE
    -- 余额不足，回滚事务
    ROLLBACK;
END IF;


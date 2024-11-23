from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class EmailRequest(_message.Message):
    __slots__ = ("id", "subject", "body", "sender", "recipients", "headers")
    class HeadersEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_FIELD_NUMBER: _ClassVar[int]
    BODY_FIELD_NUMBER: _ClassVar[int]
    SENDER_FIELD_NUMBER: _ClassVar[int]
    RECIPIENTS_FIELD_NUMBER: _ClassVar[int]
    HEADERS_FIELD_NUMBER: _ClassVar[int]
    id: str
    subject: str
    body: str
    sender: str
    recipients: _containers.RepeatedScalarFieldContainer[str]
    headers: _containers.ScalarMap[str, str]
    def __init__(self, id: _Optional[str] = ..., subject: _Optional[str] = ..., body: _Optional[str] = ..., sender: _Optional[str] = ..., recipients: _Optional[_Iterable[str]] = ..., headers: _Optional[_Mapping[str, str]] = ...) -> None: ...

class BatchEmailRequest(_message.Message):
    __slots__ = ("emails",)
    EMAILS_FIELD_NUMBER: _ClassVar[int]
    emails: _containers.RepeatedCompositeFieldContainer[EmailRequest]
    def __init__(self, emails: _Optional[_Iterable[_Union[EmailRequest, _Mapping]]] = ...) -> None: ...

class CategoryResponse(_message.Message):
    __slots__ = ("email_id", "category", "confidence", "keywords", "alternatives")
    EMAIL_ID_FIELD_NUMBER: _ClassVar[int]
    CATEGORY_FIELD_NUMBER: _ClassVar[int]
    CONFIDENCE_FIELD_NUMBER: _ClassVar[int]
    KEYWORDS_FIELD_NUMBER: _ClassVar[int]
    ALTERNATIVES_FIELD_NUMBER: _ClassVar[int]
    email_id: str
    category: str
    confidence: float
    keywords: _containers.RepeatedScalarFieldContainer[str]
    alternatives: _containers.RepeatedCompositeFieldContainer[AlternativeCategory]
    def __init__(self, email_id: _Optional[str] = ..., category: _Optional[str] = ..., confidence: _Optional[float] = ..., keywords: _Optional[_Iterable[str]] = ..., alternatives: _Optional[_Iterable[_Union[AlternativeCategory, _Mapping]]] = ...) -> None: ...

class BatchCategoryResponse(_message.Message):
    __slots__ = ("results",)
    RESULTS_FIELD_NUMBER: _ClassVar[int]
    results: _containers.RepeatedCompositeFieldContainer[CategoryResponse]
    def __init__(self, results: _Optional[_Iterable[_Union[CategoryResponse, _Mapping]]] = ...) -> None: ...

class AlternativeCategory(_message.Message):
    __slots__ = ("category", "confindence")
    CATEGORY_FIELD_NUMBER: _ClassVar[int]
    CONFINDENCE_FIELD_NUMBER: _ClassVar[int]
    category: str
    confindence: float
    def __init__(self, category: _Optional[str] = ..., confindence: _Optional[float] = ...) -> None: ...
